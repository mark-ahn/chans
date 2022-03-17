package chans

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mark-ahn/syncs"
)

type DoneNotifier interface {
	Done() <-chan struct{}
}

func Clear[T any](ch <-chan T) int {
	n := len(ch)
	for i := 0; i < n; i += 1 {
		<-ch
	}
	return n
}

type BroadCaster[T any] struct {
	srcs        map[<-chan T]struct{}
	recvs       map[chan<- T]struct{}
	add_src     chan (<-chan T)
	delete_src  chan (<-chan T)
	add_recv    chan (chan<- T)
	delete_recv chan (chan<- T)
}

func NewBroadCaster[T any](initSrcs []<-chan T, initRecvs []chan<- T) *BroadCaster[T] {
	srcs := make(map[<-chan T]struct{})
	for _, src := range initSrcs {
		srcs[src] = struct{}{}
	}
	recvs := make(map[chan<- T]struct{})
	for _, recv := range initRecvs {
		recvs[recv] = struct{}{}
	}

	return &BroadCaster[T]{
		srcs:        srcs,
		recvs:       recvs,
		add_src:     make(chan (<-chan T)),
		delete_src:  make(chan (<-chan T)),
		add_recv:    make(chan (chan<- T)),
		delete_recv: make(chan (chan<- T)),
	}
}

func (__ *BroadCaster[T]) AddSources(srcs ...<-chan T) {
	for _, src := range srcs {
		__.add_src <- src
	}
}
func (__ *BroadCaster[T]) DeleteSources(srcs ...<-chan T) {
	for _, src := range srcs {
		__.delete_src <- src
	}
}

func (__ *BroadCaster[T]) AddReceivers(recvs ...chan<- T) *BroadCaster[T] {
	for _, recv := range recvs {
		__.add_recv <- recv
	}
	return __
}
func (__ *BroadCaster[T]) DeleteReceivers(recvs ...chan<- T) *BroadCaster[T] {
	for _, recv := range recvs {
		__.delete_recv <- recv
	}
	return __
}

func (__ *BroadCaster[T]) len_read_ch() int {
	return len(__.srcs) + 4
}

func (__ *BroadCaster[T]) collect_chans(done DoneNotifier) []reflect.SelectCase {
	read_cases := make([]reflect.SelectCase, 0, __.len_read_ch()+1)
	for ch := range __.srcs {
		read_cases = append(read_cases, reflect.SelectCase{
			Chan: reflect.ValueOf(ch),
			Dir:  reflect.SelectRecv,
		})
	}
	for _, ch := range []chan (<-chan T){__.add_src, __.delete_src} {
		read_cases = append(read_cases, reflect.SelectCase{
			Chan: reflect.ValueOf(ch),
			Dir:  reflect.SelectRecv,
		})
	}
	for _, ch := range []chan (chan<- T){__.add_recv, __.delete_recv} {
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

func (__ *BroadCaster[T]) ServeThread(ctx context.Context, tctx syncs.ThreadContext) error {
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
			case <-chan T:
				d := recv.Interface().(T)
				for recv := range __.recvs {
					select {
					case recv <- d:
					default:
					}
				}
			// add/delete src
			case chan (<-chan T):
				d := recv.Interface().(<-chan T)
				switch ch {
				case __.add_src:
					__.srcs[d] = struct{}{}
				case __.delete_src:
					delete(__.srcs, d)
				}
			// add/delete recv
			case chan (chan<- T):
				d := recv.Interface().(chan<- T)
				switch ch {
				case __.add_recv:
					__.recvs[d] = struct{}{}
				case __.delete_recv:
					_, ok := __.recvs[d]
					if ok {
						delete(__.recvs, d)
						close(d)
					}
				}
			default:
				pkglog.Trace().Msgf("terminate")
				break loop
			}

		}
	}()
	return nil
}
