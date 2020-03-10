package chans

func (__ *GoChain) CaseRecvSome(ch <-chan Some, f func(v Some, ok bool) CaseControl, onEvent func(CaseResult)) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

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

func (__ *GoChain) CaseSendSome(ch chan<- Some, v Some, onEvent func(sent CaseResult), elseCh <-chan struct{}) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

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
		case <-__.ctx.Done():
			if onEvent != nil {
				onEvent(CASE_CANCEL)
			}
			break loop
		}
	}()
	return __
}

func (__ *GoChain) ConnectSome(recv <-chan Some, send chan<- Some, onEvent func(CaseResult)) *GoChain {
	__.threads.Add(1)
	go func() {
		__.threads.Done()

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
						onEvent(CASE_CLOSED)
					}
					break loop
				}
				recv_ch = nil
				send_ch = send
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
