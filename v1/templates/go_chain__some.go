package templates

import (
	"github.com/mark-ahn/chans/v1/core"
)

func (__ *GoChain) CaseRecvSome(ch <-chan Some, f func(v Some, ok bool) core.CaseControl, onEvent func(core.CaseResult)) *GoChain {
	__.AddThread(1)
	go func() {
		defer __.DoneThread()

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
			case <-__.Context().Done():
				if onEvent != nil {
					onEvent(core.CASE_CANCEL)
				}
				break loop
			}
		}

	}()
	return __
}

func (__ *GoChain) CaseSendSome(ch chan<- Some, v Some, onEvent func(sent core.CaseResult), elseCh <-chan struct{}) *GoChain {
	__.AddThread(1)
	go func() {
		defer __.DoneThread()

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
		case <-__.Context().Done():
			if onEvent != nil {
				onEvent(core.CASE_CANCEL)
			}
			break loop
		}
	}()
	return __
}

func (__ *GoChain) ConnectSome(recv <-chan Some, send chan<- Some, onEvent func(core.CaseResult)) *GoChain {
	__.AddThread(1)
	go func() {
		__.DoneThread()

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
			case <-__.Context().Done():
				if onEvent != nil {
					onEvent(core.CASE_CANCEL)
				}
				break loop
			}
		}
	}()

	return __
}
