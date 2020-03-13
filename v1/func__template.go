package chans

import (
	"context"

	"github.com/mark-ahn/syncs"
)

func _Prefix_FuncSome(ctx context.Context, f func() (Some, error), n int) <-chan Some {

	ch := make(chan Some, n)
	_Prefix_FuncSomeWith(ctx, f, ch, func() {
		close(ch)
	})
	return ch
}

func _Prefix_FuncSomeWith(ctx context.Context, f func() (Some, error), ch chan<- Some, release func()) {
	cnt := syncs.ThreadCounterFrom(ctx)

	cnt.Add(1)
	go func() {
		defer cnt.Done()
		defer func() {
			if release != nil {
				release()
			}
		}()

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
}

func _Prefix_FuncSomeSingleShot(ctx context.Context, f func() (Some, error), n int) <-chan Some {

	ch := make(chan Some, n)
	_Prefix_FuncSomeWithSingleShot(ctx, f, ch, func() {
		close(ch)
	})
	return ch
}

func _Prefix_FuncSomeWithSingleShot(ctx context.Context, f func() (Some, error), ch chan<- Some, release func()) {
	cnt := syncs.ThreadCounterFrom(ctx)

	cnt.Add(1)
	go func() {
		defer cnt.Done()
		defer func() {
			if release != nil {
				release()
			}
		}()

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
}
