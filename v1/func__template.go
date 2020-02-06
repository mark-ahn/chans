package chans

import "context"

func _Prefix_FuncSome(ctx context.Context, f func() Some, n int) <-chan Some {
	ch := make(chan Some, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()
			ch <- t
		}
	}()
	return ch
}

func _Prefix_FuncSomeSlice(ctx context.Context, f func() []Some, n int) <-chan []Some {
	ch := make(chan []Some, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()
			ch <- t
		}
	}()
	return ch
}
