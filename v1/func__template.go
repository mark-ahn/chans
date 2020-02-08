package chans

import "context"

func _Prefix_FuncSome(ctx context.Context, f func() Some, n int) <-chan Some {
	ch := make(chan Some, n)
	go func() {
		defer close(ch)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func _Prefix_FuncSomeSingleShot(ctx context.Context, f func() Some, n int) <-chan Some {
	ch := make(chan Some, n)
	go func() {
		defer close(ch)

		select {
		case <-ctx.Done():
			break
		default:
		}

		t := f()

		select {
		case <-ctx.Done():
			break
		default:
		}

		ch <- t
	}()
	return ch
}

func _Prefix_FuncSomeSlice(ctx context.Context, f func() []Some, n int) <-chan []Some {
	ch := make(chan []Some, n)
	go func() {
		defer close(ch)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func _Prefix_FuncSomeSliceSingleShot(ctx context.Context, f func() []Some, n int) <-chan []Some {
	ch := make(chan []Some, n)
	go func() {
		defer close(ch)

		select {
		case <-ctx.Done():
			break
		default:
		}

		t := f()

		select {
		case <-ctx.Done():
			break
		default:
		}

		ch <- t
	}()
	return ch
}
