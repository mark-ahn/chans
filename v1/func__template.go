package chans

import (
	"context"
	"fmt"

	"github.com/mark-ahn/chans/v1/core"
	"github.com/mark-ahn/syncs"
)

func _Prefix_FuncSome(ctx context.Context, f func() (Some, error), n int) (<-chan Some, error) {

	ch := make(chan Some, n)
	err := _Prefix_FuncSomeWith(ctx, f, ch, func() {
		close(ch)
	})
	return ch, err
}

func _Prefix_FuncSomeWith(ctx context.Context, f func() (Some, error), ch chan<- Some, release func()) error {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if release != nil {
			release()
		}
		return fmt.Errorf("cannot start thread with context which has done")
	}

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
			case core.ErrStopIter:
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
	return nil
}

func _Prefix_FuncSomeSingleShot(ctx context.Context, f func() (Some, error), n int) (<-chan Some, error) {

	ch := make(chan Some, n)
	err := _Prefix_FuncSomeWithSingleShot(ctx, f, ch, func() {
		close(ch)
	})
	return ch, err
}

func _Prefix_FuncSomeWithSingleShot(ctx context.Context, f func() (Some, error), ch chan<- Some, release func()) error {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if release != nil {
			release()
		}
		return fmt.Errorf("cannot start thread with context which has done")
	}

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
	return nil
}
