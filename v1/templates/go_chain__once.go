package templates

import (
	"context"
	"fmt"
	"reflect"

	"github.com/mark-ahn/chans/v1/core"
	"github.com/mark-ahn/syncs"
)

func CaseRecv(ctx context.Context, ch interface{}, f func(v interface{}, ok bool) core.CaseControl, onEvent func(core.CaseResult)) {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if onEvent != nil {
			onEvent(core.CASE_CANCEL)
		}
		return
	}

	go func() {
		defer cnt.Done()

		cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ctx.Done()),
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
}

func CaseSend(ctx context.Context, ch interface{}, v interface{}, f func(v interface{}, ok bool, sent core.CaseResult), elseCh interface{}) {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		f(nil, false, core.CASE_CANCEL)
	}

	go func() {
		defer cnt.Done()

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
				Chan: reflect.ValueOf(ctx.Done()),
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
}

func Connect(ctx context.Context, recv interface{}, send interface{}) error {
	cnt := syncs.ThreadCounterFrom(ctx)

	recv_v := reflect.ValueOf(recv)
	send_v := reflect.ValueOf(send)

	if recv_v.Elem().Type() != send_v.Elem().Type() {
		return fmt.Errorf("type is not matched")
	}

	ok := cnt.AddOrNot(1)
	if !ok {
		return fmt.Errorf("chans-connect: cannot start thread with context which has done")
	}

	go func() {
		defer cnt.Done()

		r_cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(recv),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ctx.Done()),
			},
		}

		w_cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(send),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ctx.Done()),
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
