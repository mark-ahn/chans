package chans_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mark-ahn/chans/v3"
)

func Test_Tokken(t *testing.T) {
	tokken := chans.NewTokken(10)
	go func() {
		fmt.Println("wait")
		<-time.After(time.Second)
		<-tokken
		fmt.Println("go")
	}()
	tokken <- struct{}{}
	fmt.Println("tokken", len(tokken))
	chans.DropAll(tokken)
	fmt.Println("drop", len(tokken))
	chans.PutFull(tokken, func(index int) struct{} { return struct{}{} })
	fmt.Println("put", len(tokken))
}
