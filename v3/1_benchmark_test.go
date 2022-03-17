package chans_test

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/mark-ahn/chans/v3"
	"github.com/mark-ahn/syncs"
)

func BenchmarkSelect(b *testing.B) {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan string, 1)

	ch_str := make(chan string, 1)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go func() {
		var tout <-chan time.Time
	loop:
		for {
			select {
			case d := <-ch1:
				ch2 <- d + 1
				tout = time.After(time.Second)
			case d := <-ch2:
				ch3 <- strconv.FormatInt(int64(d), 10)
				tout = time.After(time.Second)
			case str := <-ch3:
				ch_str <- fmt.Sprintf("[%v]", str)
				tout = nil
			case <-tout:
			case <-ctx.Done():
				break loop
			}
		}
	}()

	for i := 0; i < b.N; i += 1 {
		ch1 <- 10
		<-ch_str
	}
}

func BenchmarkSelectMany(b *testing.B) {
	const l = 10000

	chs := make([]chan int, l+1)
	for i := range chs {
		chs[i] = make(chan int, 1)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	var tout <-chan time.Time
	// for i := 0; i < l; i += 10 {
	// 	go func(s int) {
	// 	loop:
	// 		for {
	// 			select {
	// 			case d := <-chs[s+0]:
	// 				chs[s+1] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+1]:
	// 				chs[s+2] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+2]:
	// 				chs[s+3] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+3]:
	// 				chs[s+4] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+4]:
	// 				chs[s+5] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+5]:
	// 				chs[s+6] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+6]:
	// 				chs[s+7] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+7]:
	// 				chs[s+8] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+8]:
	// 				chs[s+9] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case d := <-chs[s+9]:
	// 				chs[s+10] <- d + 1
	// 				tout = time.After(time.Second)
	// 				tout = nil
	// 			case <-tout:
	// 				panic("timeout")
	// 			case <-ctx.Done():
	// 				break loop
	// 			}
	// 		}
	// 	}(i)
	// }
	for i := 0; i < l; i += 1 {
		go func(s int) {
		loop:
			for {
				select {
				case d := <-chs[s+0]:
					chs[s+1] <- d + 1
					// tout = time.After(time.Second)
					// tout = nil
				case <-tout:
					panic("timeout")
				case <-ctx.Done():
					break loop
				}
			}
		}(i)
	}

	for i := 0; i < b.N; i += 1 {
		chs[0] <- 0
		<-chs[l]
	}
}

func BenchmarkSelectModule(b *testing.B) {
	ch1 := make(chan interface{}, 1)
	ch2 := make(chan interface{}, 1)
	ch3 := make(chan interface{}, 1)

	ch_str := make(chan interface{}, 1)

	rctx, cancel := context.WithCancel(context.TODO())

	ctx, done_ch := syncs.WithThreadDoneNotify(rctx, &sync.WaitGroup{})
	defer func() {
		cancel()
		<-done_ch
	}()

	chans.CaseRecv(ctx, ch1, func(recv interface{}, ok bool) chans.CaseControl {
		d, ok := recv.(int)

		chans.CaseSend(ctx, ch2, any(d+1), func(sent chans.CaseResult) {}, nil)
		return chans.CASE_OK
	}, nil)
	chans.CaseRecv(ctx, ch2, func(recv interface{}, ok bool) chans.CaseControl {
		d, ok := recv.(int)

		chans.CaseSend(ctx, ch3, any(strconv.FormatInt(int64(d), 10)), func(sent chans.CaseResult) {}, nil)
		return chans.CASE_OK
	}, nil)
	chans.CaseRecv(ctx, ch3, func(recv interface{}, ok bool) chans.CaseControl {
		str, ok := recv.(string)
		if !ok {
			return chans.CASE_OK
		}
		chans.CaseSend(ctx, ch_str, any(fmt.Sprintf("[%v]", str)), func(sent chans.CaseResult) {}, nil)

		return chans.CASE_OK
	}, nil)

	for i := 0; i < b.N; i += 1 {
		ch1 <- 10
		<-ch_str
	}

}

func BenchmarkSelectModuleMany(b *testing.B) {
	const l = 10000
	chs := make([]chan interface{}, l+1)
	for i := range chs {
		chs[i] = make(chan interface{}, 1)
	}
	rctx, cancel := context.WithCancel(context.TODO())
	ctx, done_ch := syncs.WithThreadDoneNotify(rctx, &sync.WaitGroup{})

	for i := 0; i < l; i += 1 {
		func(i int) {
			chans.CaseRecv(ctx, chs[i], func(recv interface{}, ok bool) chans.CaseControl {
				d, ok := recv.(int)
				chans.CaseSend(ctx, chs[i+1], any(d+1), func(sent chans.CaseResult) {
				}, nil)
				return chans.CASE_OK
			}, nil)
		}(i)
	}
	defer func() {
		cancel()
		<-done_ch
	}()

	for i := 0; i < b.N; i += 1 {
		chs[0] <- 0
		<-chs[l]
	}
}

func BenchmarkSelectModuleManyWithType(b *testing.B) {
	const l = 10000
	chs := make([]chan int, l+1)
	for i := range chs {
		chs[i] = make(chan int, 1)
	}
	rctx, cancel := context.WithCancel(context.TODO())
	ctx, done_ch := syncs.WithThreadDoneNotify(rctx, &sync.WaitGroup{})

	for i := 0; i < l; i += 1 {
		func(i int) {
			chans.CaseRecv(ctx, chs[i], func(d int, ok bool) chans.CaseControl {
				chans.CaseSend(ctx, chs[i+1], d+1, nil, nil)
				return chans.CASE_OK
			}, nil)
		}(i)
	}
	defer func() {
		cancel()
		<-done_ch
	}()

	for i := 0; i < b.N; i += 1 {
		chs[0] <- 0
		<-chs[l]
	}
}
