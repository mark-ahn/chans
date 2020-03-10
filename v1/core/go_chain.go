package core

import (
	"context"
	"fmt"
	"sync"
)

type CaseControl int

const (
	CASE_OK CaseControl = iota
	CASE_BLOCK
	CASE_SELECTION_TEARDOWN
	CASE_DO_NOTHING
)

type CaseResult int

const (
	CASE_SENT CaseResult = iota
	CASE_ELSE
	CASE_CANCEL
	CASE_CLOSED
	CASE_STOP
)

var ErrStopMap = fmt.Errorf("Stop Map")
var ErrSkipMap = fmt.Errorf("Skip Map")

type Chainable interface {
	Context() context.Context
	AddThread(int)
	DoneThread()
}

type GoChain struct {
	ctx context.Context

	threads sync.WaitGroup
	doneCh  chan struct{}
}

func (__ *GoChain) DoneNotify() <-chan struct{} {
	return __.doneCh
}

func (__ *GoChain) Context() context.Context {
	return __.ctx
}

func (__ *GoChain) AddThread(i int) {
	__.threads.Add(i)
}

func (__ *GoChain) DoneThread() {
	__.threads.Done()
}

func NewGoChain(ctx context.Context) *GoChain {
	__ := &GoChain{
		ctx: ctx,

		doneCh: make(chan struct{}),
	}

	go func() {
		defer close(__.doneCh)
		defer __.threads.Wait()

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			}
		}
	}()

	return __
}