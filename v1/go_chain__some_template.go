package chans

func (__ *GoChain) CaseRecvSome(ch <-chan Some, f func(v Some, ok bool) CaseControl) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

	loop:
		for {
			select {
			case v, ok := <-ch:
				switch f(v, ok) {
				case CASE_OK:
					continue
				default:
					break loop
				}
			case <-__.ctx.Done():
				break loop
			}
		}

	}()
	return __
}

func (__ *GoChain) CaseSendSome(ch chan<- Some, v Some, f func(sent CaseSend), elseCh <-chan struct{}) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

	loop:
		select {
		case ch <- v:
			if f != nil {
				f(CASE_SENT)
			}
		case <-elseCh:
			if f != nil {
				f(CASE_ELSE)
			}
		case <-__.ctx.Done():
			if f != nil {
				f(CASE_CANCEL)
			}
			break loop
		}
	}()
	return __
}

func (__ *GoChain) ConnectSome(recv <-chan Some, send chan<- Some) *GoChain {
	__.threads.Add(1)
	go func() {
		__.threads.Done()

	loop:
		for {
			select {
			case d, ok := <-recv:
				if !ok {
					break loop
				}
				send <- d
			case <-__.ctx.Done():
				break loop
			}
		}
	}()

	return __
}
