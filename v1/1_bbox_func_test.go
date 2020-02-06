package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mark-ahn/chans/v1"
)

func TestFunc(t *testing.T) {
	ch := chans.OfFuncByteSlice(context.Background(), func() []byte {
		<-time.After(200 * time.Millisecond)
		return []byte{}
	}, 1)
	for _ = range []int{1, 2, 3, 4, 5} {
		b := <-ch
		fmt.Println(b)
	}
}
