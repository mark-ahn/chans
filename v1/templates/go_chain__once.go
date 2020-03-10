package templates

import (
	"fmt"
	"reflect"

	"github.com/mark-ahn/chans/v1/core"
)

type GoChain struct{ core.Chainable }

func WithGoChain(chain core.Chainable) *GoChain {
	return &GoChain{Chainable: chain}
}

func (__ *GoChain) CaseRecv(ch interface{}, f func(v interface{}, ok bool) core.CaseControl, onEvent func(core.CaseResult)) *GoChain {
	__.AddThread(1)
	go func() {
		defer __.DoneThread()

		cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.Context().Done()),
			},
		}

	loop:
		for {

			i, v, ok := reflect.Select(cases)

			switch i {
			case 0:
				switch f(v.Interface(), ok) {
				case core.CASE_OK:
					continue
				default:
					code := core.CASE_STOP
					switch ok {
					case false:
						code = core.CASE_CLOSED
					}
					if onEvent != nil {
						onEvent(code)
					}
					break loop
				}
			default:
				if onEvent != nil {
					onEvent(core.CASE_CANCEL)
				}
				break loop
			}
		}
	}()
	return __
}

func (__ *GoChain) CaseSend(ch interface{}, v interface{}, f func(v interface{}, ok bool, sent core.CaseResult), elseCh interface{}) *GoChain {
	__.AddThread(1)
	go func() {
		defer __.DoneThread()

		cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(ch),
				Send: reflect.ValueOf(v),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(elseCh),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.Context().Done()),
			},
		}

		i, v, ok := reflect.Select(cases)
		switch i {
		case 0:
			f(nil, false, core.CASE_SENT)
		case 1:
			f(v.Interface(), ok, core.CASE_ELSE)
		default:
			f(nil, false, core.CASE_CANCEL)
		}

	}()
	return __
}

func (__ *GoChain) Connect(recv interface{}, send interface{}) error {
	recv_v := reflect.ValueOf(recv)
	send_v := reflect.ValueOf(send)

	if recv_v.Elem().Type() != send_v.Elem().Type() {
		return fmt.Errorf("type is not matched")
	}

	__.AddThread(1)
	go func() {
		defer __.DoneThread()

		r_cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(recv),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.Context().Done()),
			},
		}

		w_cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(send),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.Context().Done()),
			},
		}

	loop:
		for {
			i, v, ok := reflect.Select(r_cases)
			if !(i == 0 && ok) {
				break loop
			}
			w_cases[0].Send = v
			i, _, _ = reflect.Select(w_cases)
			if !(i == 0) {
				break loop
			}
		}
	}()

	return nil
}
