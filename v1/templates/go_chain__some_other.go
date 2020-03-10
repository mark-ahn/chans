package templates

import "github.com/mark-ahn/chans/v1/core"

func (__ *GoChain) MapSomeToOther(recv <-chan Some, send chan<- Other, mapF func(Some, bool) (Other, error), onEvent func(core.CaseResult)) *GoChain {
	__.AddThread(1)
	go func() {
		__.DoneThread()

		var err error
		var send_ch chan<- Other
		var recv_ch <-chan Some = recv
		var to_send Other

	loop:
		for {
			select {
			case d, ok := <-recv_ch:
				to_send, err = mapF(d, ok)
				switch err {
				case core.ErrSkipMap:
					continue
				case nil:
					recv_ch = nil
					send_ch = send
				default:
					if onEvent != nil {
						switch ok {
						case false:
							onEvent(core.CASE_CLOSED)
						default:
							onEvent(core.CASE_STOP)
						}
					}
					break loop
				}

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

func (__ *GoChain) CaseSendSomeOrOther(ch chan<- Some, v Some, onEvent func(sent core.CaseResult), elseCh <-chan Other, elseF func(v Other, ok bool)) *GoChain {
	__.AddThread(1)
	go func() {
		defer __.DoneThread()

	loop:
		select {
		case ch <- v:
			if onEvent != nil {
				onEvent(core.CASE_SENT)
			}
		case v, ok := <-elseCh:
			if elseF != nil {
				elseF(v, ok)
			}
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
