package chans

func (__ *GoChain) MapSomeToOther(recv <-chan Some, send chan<- Other, mapF func(Some, bool) (Other, error), onEvent func(CaseResult)) *GoChain {
	__.threads.Add(1)
	go func() {
		__.threads.Done()

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
				case ErrSkipMap:
					continue
				case nil:
					recv_ch = nil
					send_ch = send
				default:
					if onEvent != nil {
						switch ok {
						case false:
							onEvent(CASE_CLOSED)
						default:
							onEvent(CASE_STOP)
						}
					}
					break loop
				}

			case send_ch <- to_send:
				send_ch = nil
				recv_ch = recv

			case <-__.ctx.Done():
				if onEvent != nil {
					onEvent(CASE_CANCEL)
				}
				break loop
			}
		}
	}()

	return __
}

func (__ *GoChain) CaseSendSomeOrOther(ch chan<- Some, v Some, onEvent func(sent CaseResult), elseCh <-chan Other, elseF func(v Other, ok bool)) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

	loop:
		select {
		case ch <- v:
			if onEvent != nil {
				onEvent(CASE_SENT)
			}
		case v, ok := <-elseCh:
			if elseF != nil {
				elseF(v, ok)
			}
			if onEvent != nil {
				onEvent(CASE_ELSE)
			}
		case <-__.ctx.Done():
			if onEvent != nil {
				onEvent(CASE_CANCEL)
			}
			break loop
		}
	}()
	return __
}
