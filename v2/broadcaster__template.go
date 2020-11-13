package chans

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mark-ahn/syncs"
)

func _Prefix_SomeClear(ch <-chan Some) int {
	n := len(ch)
	for i := 0; i < n; i += 1 {
		<-ch
	}
	return n
}

type _Prefix_SomeBroadCaster struct {
	srcs        map[<-chan Some]struct{}
	recvs       map[chan<- Some]struct{}
	add_src     chan (<-chan Some)
	delete_src  chan (<-chan Some)
	add_recv    chan (chan<- Some)
	delete_recv chan (chan<- Some)
}

func New_Prefix_SomeBroadCaster(initSrcs []<-chan Some, initRecvs []chan<- Some) *_Prefix_SomeBroadCaster {
	// fmt.Println("New_Prefix_SomeBroadCaster")
	srcs := make(map[<-chan Some]struct{})
	for _, src := range initSrcs {
		srcs[src] = struct{}{}
	}
	recvs := make(map[chan<- Some]struct{})
	for _, recv := range initRecvs {
		recvs[recv] = struct{}{}
	}

	return &_Prefix_SomeBroadCaster{
		srcs:        srcs,
		recvs:       recvs,
		add_src:     make(chan (<-chan Some)),
		delete_src:  make(chan (<-chan Some)),
		add_recv:    make(chan (chan<- Some)),
		delete_recv: make(chan (chan<- Some)),
	}
}

func (__ *_Prefix_SomeBroadCaster) AddSources(srcs ...<-chan Some) {
	// fmt.Println("AddSources")
	for _, src := range srcs {
		__.add_src <- src
	}
}
func (__ *_Prefix_SomeBroadCaster) DeleteSources(srcs ...<-chan Some) {
	// fmt.Println("DeleteSources")
	for _, src := range srcs {
		__.delete_src <- src
	}
}

func (__ *_Prefix_SomeBroadCaster) AddReceivers(recvs ...chan<- Some) *_Prefix_SomeBroadCaster {
	for _, recv := range recvs {
		__.add_recv <- recv
	}
	return __
}
func (__ *_Prefix_SomeBroadCaster) DeleteReceivers(recvs ...chan<- Some) *_Prefix_SomeBroadCaster {
	for _, recv := range recvs {
		__.delete_recv <- recv
	}
	return __
}

func (__ *_Prefix_SomeBroadCaster) len_read_ch() int {
	return len(__.srcs) + 4
}

func (__ *_Prefix_SomeBroadCaster) collect_chans(done DoneNotifier) []reflect.SelectCase {
	read_cases := make([]reflect.SelectCase, 0, __.len_read_ch()+1)
	for ch := range __.srcs {
		// fmt.Println(ch)
		read_cases = append(read_cases, reflect.SelectCase{
			Chan: reflect.ValueOf(ch),
			Dir:  reflect.SelectRecv,
		})
	}
	for _, ch := range []chan (<-chan Some){__.add_src, __.delete_src} {
		read_cases = append(read_cases, reflect.SelectCase{
			Chan: reflect.ValueOf(ch),
			Dir:  reflect.SelectRecv,
		})
	}
	for _, ch := range []chan (chan<- Some){__.add_recv, __.delete_recv} {
		read_cases = append(read_cases, reflect.SelectCase{
			Chan: reflect.ValueOf(ch),
			Dir:  reflect.SelectRecv,
		})
	}

	read_cases = append(read_cases, reflect.SelectCase{
		Chan: reflect.ValueOf(done.Done()),
		Dir:  reflect.SelectRecv,
	})

	return read_cases
}

func (__ *_Prefix_SomeBroadCaster) ServeThread(ctx context.Context, tctx syncs.ThreadContext) error {
	// fmt.Println("ServeThread")
	th_cnt := syncs.ThreadCounterFrom(ctx)
	ok := th_cnt.AddOrNot(1)
	if !ok {
		return fmt.Errorf("cannot serve thread: context done")
	}

	go func() {
		defer th_cnt.Done()
		defer func() {
			// fmt.Println("defer")
			for ch := range __.recvs {
				close(ch)
			}
		}()
	loop:
		for {
			cases := __.collect_chans(ctx)
			// fmt.Println(cases)
			chosen, recv, recv_ok := reflect.Select(cases)
			if !recv_ok {
				tctx.Break(fmt.Errorf("receive channel #%v is broken", chosen))
				break loop
			}
			// fmt.Println(chosen, recv, recv_ok)
			switch ch := cases[chosen].Chan.Interface().(type) {
			// srcs
			case <-chan Some:
				d := recv.Interface().(Some)
				for recv := range __.recvs {
					select {
					case recv <- d:
					default:
					}
				}
			// add/delete src
			case chan (<-chan Some):
				d := recv.Interface().(<-chan Some)
				switch ch {
				case __.add_src:
					__.srcs[d] = struct{}{}
				case __.delete_src:
					delete(__.srcs, d)
				}
				// add/delete recv
			case chan (chan<- Some):
				d := recv.Interface().(chan<- Some)
				switch ch {
				case __.add_recv:
					__.recvs[d] = struct{}{}
				case __.delete_recv:
					delete(__.recvs, d)
					close(d)
				}
			default:
				fmt.Println("terminate")
				break loop
			}

		}
	}()
	return nil
}
