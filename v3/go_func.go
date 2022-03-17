package chans

import (
	"context"
	"fmt"

	"github.com/mark-ahn/syncs"
)

func MakePushWith[T any](ctx context.Context, f func() (T, error), n int) (<-chan T, error) {

	ch := make(chan T, n)
	err := Push(ctx, ch, f, func() {
		close(ch)
	})
	return ch, err
}

// push every value returned from f to channel
func Push[T any](ctx context.Context, ch chan<- T, f func() (T, error), release func()) error {
	cnt := syncs.ThreadCounterFrom(ctx)

	ok := cnt.AddOrNot(1)
	if !ok {
		if release != nil {
			release()
		}
		return fmt.Errorf("cannot start thread cause context done")
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
	return nil
}

func MakePushSingleShotWith[T any](ctx context.Context, f func() (T, error), n int) (<-chan T, error) {
	ch := make(chan T, n)
	err := PushSingleShot(ctx, f, ch, func() {
		close(ch)
	})
	return ch, err
}

func PushSingleShot[T any](ctx context.Context, f func() (T, error), ch chan<- T, release func()) error {
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
