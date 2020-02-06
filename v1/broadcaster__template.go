package chans

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mark-ahn/hexa"
)

func _Prefix_SomeClear(ch <-chan Some) int {
	n := len(ch)
	for i := 0; i < n; i += 1 {
		<-ch
	}
	return n
}

type _Prefix_SomeBroadCaster struct {
	in  []<-chan Some
	out []chan<- Some
}

func New_Prefix_SomeBroadCaster() *_Prefix_SomeBroadCaster {
	return &_Prefix_SomeBroadCaster{
		in:  make([]<-chan Some, 0),
		out: make([]chan<- Some, 0),
	}
}

func (__ *_Prefix_SomeBroadCaster) AddSources(ins ...<-chan Some) *_Prefix_SomeBroadCaster {
	__.in = append(__.in, ins...)
	return __
}

func (__ *_Prefix_SomeBroadCaster) AddReceivers(outs ...chan<- Some) *_Prefix_SomeBroadCaster {
	__.out = append(__.out, outs...)
	return __
}

func (__ *_Prefix_SomeBroadCaster) Serve() hexa.StoppableOne {
	ctx_stop := hexa.NewContextStop(context.Background())
	go func() {
		defer func() {
			for _, ch := range __.out {
				close(ch)
			}
			ctx_stop.InClose()
		}()

		read_cases := make([]reflect.SelectCase, len(__.in)+1)
		for i := range __.in {
			read_cases[i] = reflect.SelectCase{
				Chan: reflect.ValueOf(__.in[i]),
				Dir:  reflect.SelectRecv,
			}
		}
		read_cases[len(__.in)] = reflect.SelectCase{
			Chan: reflect.ValueOf(ctx_stop.InDoneNotify()),
			Dir:  reflect.SelectRecv,
		}
	loop:
		for {
			chosen, recv, recvOK := reflect.Select(read_cases)
			switch {
			case chosen < len(__.in):
			default:
				break loop
			}
			if !recvOK {
				ctx_stop.InBreak(fmt.Errorf("receive channel #%v is broken", chosen))
				continue
			}
			d, ok := recv.Interface().(Some)
			if !ok {
				ctx_stop.InBreak(fmt.Errorf("receive channel #%v is broken", chosen))
				continue
			}
			for i := range __.out {
				select {
				case __.out[i] <- d:
				default:
					fmt.Printf("send channel #%v is blocked - skip sending %v", i, d)
				}
			}
		}
	}()
	return ctx_stop
}
