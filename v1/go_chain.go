package chans

//go:generate genny -in go_chain__some_template.go -out go_chain__some_template__gen.go gen "_Prefix_=Of Some=Bytes,BUILTINS,interface{},struct{}"
//go:generate genny -in go_chain__some_other_template.go -out go_chain__some_other_template__gen.go gen "_Prefix_=Of Some=Bytes,BUILTINS,interface{},struct{} Other=Bytes,BUILTINS,interface{},struct{},time.Time"

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

type CaseControl int

const (
	CASE_OK CaseControl = iota
	CASE_BLOCK
	CASE_SELECTION_TEARDOWN
	CASE_DO_NOTHING
)

type CaseResult int

const (
	CASE_SENT CaseResult = iota
	CASE_ELSE
	CASE_CANCEL
	CASE_CLOSED
	CASE_STOP
)

var ErrStopMap = fmt.Errorf("Stop Map")
var ErrSkipMap = fmt.Errorf("Skip Map")

type GoChain struct {
	ctx context.Context

	threads sync.WaitGroup
	doneCh  chan struct{}
}

func NewGoChain(ctx context.Context) *GoChain {
	__ := &GoChain{
		ctx: ctx,

		doneCh: make(chan struct{}),
	}

	go func() {
		defer close(__.doneCh)
		defer __.threads.Wait()

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			}
		}
	}()

	return __
}

func (__ *GoChain) CaseRecv(ch interface{}, f func(v interface{}, ok bool) CaseControl, onEvent func(CaseResult)) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

		cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.ctx.Done()),
			},
		}

	loop:
		for {

			i, v, ok := reflect.Select(cases)

			switch i {
			case 0:
				switch f(v.Interface(), ok) {
				case CASE_OK:
					continue
				default:
					code := CASE_STOP
					switch ok {
					case false:
						code = CASE_CLOSED
					}
					if onEvent != nil {
						onEvent(code)
					}
					break loop
				}
			default:
				onEvent(CASE_CANCEL)
				break loop
			}
		}
	}()
	return __
}

func (__ *GoChain) CaseSend(ch interface{}, v interface{}, f func(v interface{}, ok bool, sent CaseResult), elseCh interface{}) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

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
				Chan: reflect.ValueOf(__.ctx.Done()),
			},
		}

		i, v, ok := reflect.Select(cases)
		switch i {
		case 0:
			f(nil, false, CASE_SENT)
		case 1:
			f(v.Interface(), ok, CASE_ELSE)
		default:
			f(nil, false, CASE_CANCEL)
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

	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

		r_cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectSend,
				Chan: reflect.ValueOf(recv),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.ctx.Done()),
			},
		}

		w_cases := []reflect.SelectCase{
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(send),
			},
			{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(__.ctx.Done()),
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

func (__ *GoChain) DoneNotify() <-chan struct{} {
	return __.doneCh
}
