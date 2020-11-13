package chans_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mark-ahn/chans/v2"
	"github.com/mark-ahn/syncs"
)

func TestIntCaster(t *testing.T) {
	src_one := make(chan int, 1)
	src_two := make(chan int, 1)

	dest_one := make(chan int, 1)
	dest_two := make(chan int, 1)
	dest_three := make(chan int, 1)

	caster := chans.NewOfIntBroadCaster(nil, nil)
	// ctx, stop := context.WithCancel(context.TODO())
	ctx := context.TODO()
	cast_srv, err := syncs.Serve(ctx, caster)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		// stop()
		cast_srv.Break(nil)
		fmt.Println("stop")
		<-cast_srv.Done()
		fmt.Println("exit")
	}()
	caster.AddSources(src_one, src_two)
	caster.AddReceivers(dest_one, dest_two, dest_three)
	time.Sleep(time.Second)

	src_one <- 10
	for _, ch := range []chan int{
		dest_one, dest_two, dest_three,
	} {
		recv := <-ch
		if recv != 10 {
			t.Errorf("expect %v, got %v", 10, recv)
		}
	}

	src_two <- 20
	for _, ch := range []chan int{
		dest_one, dest_two, dest_three,
	} {
		recv := <-ch
		if recv != 20 {
			t.Errorf("expect %v, got %v", 10, recv)
		}
	}
	fmt.Printf("done \n")
}
