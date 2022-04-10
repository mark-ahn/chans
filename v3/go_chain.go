package chans

import (
	"context"
	"fmt"

	"github.com/mark-ahn/syncs"
)

// func Connect[T any](ctx context.Context, recv <-chan T, send chan<- T) error {
// 	cnt := syncs.ThreadCounterFrom(ctx)

// 	ok := cnt.AddOrNot(1)
// 	if !ok {
// 		return fmt.Errorf("chans-connect: cannot start thread cause context done")
// 	}

// 	go func() {
// 		defer cnt.Done()

// 	loop:
// 		for {
// 			var ok bool
// 			var v T

// 			select {
// 			case v, ok = <-recv:
// 			case <-ctx.Done():
// 				break loop
// 			}
// 			if !ok {
// 				break loop
// 			}

// 			select {
// 			case send <- v:
// 			case <-ctx.Done():
// 				break loop
// 			}
// 		}
// 	}()

// 	return nil
// }

func bypass[T any](t T, b bool) (T, error) {
	if !b {
		return t, ErrStopIter
	}
	return t, nil
}
func Connect[T any](ctx context.Context, recv <-chan T, send chan<- T, onEvent func(CaseResult)) error {
	return ConnectFunc(ctx, recv, send, bypass[T], onEvent)
}
func ConnectFunc[T any, U any](ctx context.Context, recv <-chan T, send chan<- U, mapF func(T, bool) (U, error), onEvent func(CaseResult)) error {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if onEvent != nil {
			onEvent(CASE_CANCEL)
		}
		return fmt.Errorf("cannot start thread cause context done")
	}

	go func() {
		defer cnt.Done()

		var err error
		var send_ch chan<- U
		var recv_ch <-chan T = recv
		var to_send U

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

			case <-ctx.Done():
				if onEvent != nil {
					onEvent(CASE_CANCEL)
				}
				break loop
			}
		}
	}()

	return nil
}
