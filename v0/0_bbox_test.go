package chans_test

import (
	"fmt"
	"testing"

	"github.com/mark-ahn/chans/v0"
)

func TestIntCaster(t *testing.T) {
	src_one := make(chan int, 1)
	src_two := make(chan int, 1)

	dest_one := make(chan int, 1)
	dest_two := make(chan int, 1)
	dest_three := make(chan int, 1)

	cast_srv := chans.NewOfIntBroadCaster().AddSources(src_one, src_two).AddReceivers(dest_one, dest_two, dest_three).Serve()
	defer func() {
		cast_srv.Close()
		<-cast_srv.DoneNotify()
	}()

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
