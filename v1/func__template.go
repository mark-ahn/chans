package chans

import (
	"context"

	"github.com/mark-ahn/syncs"
)

func _Prefix_FuncSome(ctx context.Context, f func() (Some, error), n int) <-chan Some {
	cnt := syncs.ThreadCounterFrom(ctx)

	ch := make(chan Some, n)
	cnt.Add(1)
	go func() {
		defer cnt.Done()
		defer close(ch)

	loop:
		for {
			t, err := f()
			switch err {
			case ErrStopIter:
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
	cnt := syncs.ThreadCounterFrom(ctx)

	ch := make(chan Some, n)
	cnt.Add(1)
	go func() {
		defer cnt.Done()
		defer close(ch)

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
