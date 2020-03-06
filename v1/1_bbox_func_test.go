package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mark-ahn/chans/v1"
)

func TestFunc(t *testing.T) {
	ch := chans.OfFuncBytes(context.Background(), func() ([]byte, error) {
		<-time.After(200 * time.Millisecond)
		return []byte{}, nil
	}, 1)
	for _ = range []int{1, 2, 3, 4, 5} {
		b, ok := <-ch
		fmt.Println(b, ok)
	}
}

func TestFuncSingleShot(t *testing.T) {
	ch := chans.OfFuncBytesSingleShot(context.Background(), func() ([]byte, error) {
		<-time.After(200 * time.Millisecond)
		return []byte{}, nil
	}, 1)
	for _ = range []int{1, 2, 3, 4, 5} {
		b, ok := <-ch
		fmt.Println(b, ok)
	}
}
