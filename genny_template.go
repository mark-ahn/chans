package chans

import (
	"context"
	"fmt"
	"reflect"

	"github.com/cheekybits/genny/generic"
	"github.com/mark-ahn/hexa"
)

type Some generic.Type

type _Prefix_SomeMultiCaster struct {
	in  []<-chan Some
	out []chan<- Some
}

func New_Prefix_SomeMultiCaster() *_Prefix_SomeMultiCaster {
	return &_Prefix_SomeMultiCaster{
		in:  make([]<-chan Some, 0),
		out: make([]chan<- Some, 0),
	}
}

func (__ *_Prefix_SomeMultiCaster) AddSources(ins ...<-chan Some) *_Prefix_SomeMultiCaster {
	__.in = append(__.in, ins...)
	return __
}

func (__ *_Prefix_SomeMultiCaster) Add(outs ...chan<- Some) *_Prefix_SomeMultiCaster {
	__.out = append(__.out, outs...)
	return __
}

func (__ *_Prefix_SomeMultiCaster) Serve() hexa.StoppableOne {
	dctx := hexa.NewContextStop(context.Background())
	go func() {
		defer func() {
			for _, ch := range __.out {
				close(ch)
			}
			dctx.InClose()
		}()

		read_cases := make([]reflect.SelectCase, len(__.in)+1)
		for i := range __.in {
			read_cases[i] = reflect.SelectCase{
				Chan: reflect.ValueOf(__.in[i]),
				Dir:  reflect.SelectRecv,
			}
		}
		read_cases[len(__.in)] = reflect.SelectCase{
			Chan: reflect.ValueOf(dctx.InDoneNotify()),
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
				dctx.InBreak(fmt.Errorf("receive channel #%v is broken", chosen))
				continue
			}
			d, ok := recv.Interface().(Some)
			if !ok {
				dctx.InBreak(fmt.Errorf("receive channel #%v is broken", chosen))
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
	return dctx
}
