package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mark-ahn/chans/v3"
)

func TestFunc(t *testing.T) {
	ch, err := chans.MakePushWith(context.Background(), func() ([]byte, error) {
		<-time.After(200 * time.Millisecond)
		return []byte{}, nil
	}, 1)
	if err != nil {
		t.Fatal(err)
	}
	for range []int{1, 2, 3, 4, 5} {
		b, ok := <-ch
		fmt.Println(b, ok)
	}
}

func TestFuncSingleShot(t *testing.T) {
	ch, err := chans.MakePushSingleShotWith(context.Background(), func() ([]byte, error) {
		<-time.After(200 * time.Millisecond)
		return []byte{}, nil
	}, 1)
	if err != nil {
		t.Fatal(err)
	}
	for range []int{1, 2, 3, 4, 5} {
		b, ok := <-ch
		fmt.Println(b, ok)
	}
}
