package chans

import (
	"context"

	"github.com/mark-ahn/syncs"
)

func CaseRecv[Some any](ctx context.Context, ch <-chan Some, f func(v Some, ok bool) CaseControl, onEvent func(CaseResult)) {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if onEvent != nil {
			onEvent(CASE_CANCEL)
		}
		return
	}

	go func() {
		defer cnt.Done()

	loop:
		for {
			select {
			case v, ok := <-ch:
				res := f(v, ok)
				switch {
				case ok && res == CASE_OK:
					continue
				default:
					if onEvent != nil {
						onEvent(CASE_CLOSED)
					}
					break loop
				}
			case <-ctx.Done():
				if onEvent != nil {
					onEvent(CASE_CANCEL)
				}
				break loop
			}
		}

	}()
}

func CaseSend[Some any](ctx context.Context, ch chan<- Some, v Some, onEvent func(sent CaseResult), elseCh <-chan struct{}) {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if onEvent != nil {
			onEvent(CASE_CANCEL)
		}
		return
	}

	go func() {
		defer cnt.Done()

	loop:
		select {
		case ch <- v:
			if onEvent != nil {
				onEvent(CASE_SENT)
			}
		case <-elseCh:
			if onEvent != nil {
				onEvent(CASE_ELSE)
			}
		case <-ctx.Done():
			if onEvent != nil {
				onEvent(CASE_CANCEL)
			}
			break loop
		}
	}()
}
