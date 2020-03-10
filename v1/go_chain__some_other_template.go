package chans

func (__ *GoChain) MapSomeToOther(recv <-chan Some, send chan<- Other, f func(Some, bool) (Other, error)) *GoChain {
	__.threads.Add(1)
	go func() {
		__.threads.Done()

	loop:
		for {
			select {
			case d, ok := <-recv:
				o, err := f(d, ok)
				switch err {
				case ErrSkipMap:
					continue
				case nil:
					send <- o
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

func (__ *GoChain) CaseSendSomeOrOther(ch chan<- Some, v Some, f func(sent CaseSend, v Other, ok bool), elseCh <-chan Other) *GoChain {
	__.threads.Add(1)
	go func() {
		defer __.threads.Done()

	loop:
		select {
		case ch <- v:
			var d Other
			if f != nil {
				f(CASE_SENT, d, false)
			}
		case v, ok := <-elseCh:
			if f != nil {
				f(CASE_ELSE, v, ok)
			}
		case <-__.ctx.Done():
			var d Other
			if f != nil {
				f(CASE_CANCEL, d, false)
			}
			break loop
		}
	}()
	return __
}
