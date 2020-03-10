package chans

import (
	"context"

	"github.com/mark-ahn/chans/v1/core"
)

func _Prefix_FuncSome(ctx context.Context, f func() (Some, error), n int) <-chan Some {
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

			t, err := f()
			switch err {
			case core.StopIterationError:
				break loop
			case nil:
			default:
				continue
			}

			select {
			case <-ctx.Done():
				break loop
			case ch <- t:
			}

		}
	}()
	return ch
}

func _Prefix_FuncSomeSingleShot(ctx context.Context, f func() (Some, error), n int) <-chan Some {
	ch := make(chan Some, n)
	go func() {
		defer close(ch)

		select {
		case <-ctx.Done():
			return
		default:
		}

		t, err := f()
		if err != nil {
			return
		}

		select {
		case <-ctx.Done():
			return
		case ch <- t:
		}

	}()
	return ch
}
