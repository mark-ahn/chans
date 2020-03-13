package templates

import (
	"context"

	"github.com/mark-ahn/chans/v1/core"
	"github.com/mark-ahn/syncs"
)

func CaseRecvSome(ctx context.Context, ch <-chan Some, f func(v Some, ok bool) core.CaseControl, onEvent func(core.CaseResult)) {
	cnt := syncs.ThreadCounterFrom(ctx)

	cnt.Add(1)
	go func() {
		defer cnt.Done()

	loop:
		for {
			select {
			case v, ok := <-ch:
				res := f(v, ok)
				switch {
				case ok && res == core.CASE_OK:
					continue
				default:
					if onEvent != nil {
						onEvent(core.CASE_CLOSED)
					}
					break loop
				}
			case <-ctx.Done():
				if onEvent != nil {
					onEvent(core.CASE_CANCEL)
				}
				break loop
			}
		}

	}()
}

func CaseSendSome(ctx context.Context, ch chan<- Some, v Some, onEvent func(sent core.CaseResult), elseCh <-chan struct{}) {
	cnt := syncs.ThreadCounterFrom(ctx)

	cnt.Add(1)
	go func() {
		defer cnt.Done()

	loop:
		select {
		case ch <- v:
			if onEvent != nil {
				onEvent(core.CASE_SENT)
			}
		case <-elseCh:
			if onEvent != nil {
				onEvent(core.CASE_ELSE)
			}
		case <-ctx.Done():
			if onEvent != nil {
				onEvent(core.CASE_CANCEL)
			}
			break loop
		}
	}()
}

func ConnectSome(ctx context.Context, recv <-chan Some, send chan<- Some, onEvent func(core.CaseResult)) {
	cnt := syncs.ThreadCounterFrom(ctx)

	cnt.Add(1)
	go func() {
		defer cnt.Done()

		var ok bool
		var recv_ch <-chan Some = recv
		var send_ch chan<- Some
		var to_send Some

	loop:
		for {
			select {
			case to_send, ok = <-recv_ch:
				if !ok {
					if onEvent != nil {
						onEvent(core.CASE_CLOSED)
					}
					break loop
				}
				recv_ch = nil
				send_ch = send
			case send_ch <- to_send:
				send_ch = nil
				recv_ch = recv
			case <-ctx.Done():
				if onEvent != nil {
					onEvent(core.CASE_CANCEL)
				}
				break loop
			}
		}
	}()

}
