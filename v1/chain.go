package chans

import (
	"context"

	"github.com/mark-ahn/chans/v1/core"
	"github.com/mark-ahn/chans/v1/generated"
)

type Chainable = core.Chainable

var ErrSkipMap = core.ErrSkipMap
var ErrStopMap = core.ErrStopMap

type Chain = generated.Chain

func WithChain(chain core.Chainable) *Chain {
	return generated.WithChain(chain)
}

func NewGoChain(ctx context.Context) *core.GoChain {
	return core.NewGoChain(ctx)
}

type CaseControl = core.CaseControl

const (
	CASE_OK                 = core.CASE_OK
	CASE_BLOCK              = core.CASE_BLOCK
	CASE_SELECTION_TEARDOWN = core.CASE_SELECTION_TEARDOWN
	CASE_DO_NOTHING         = core.CASE_DO_NOTHING
)

type CaseResult = core.CaseResult

const (
	CASE_SENT   = core.CASE_SENT
	CASE_ELSE   = core.CASE_ELSE
	CASE_CANCEL = core.CASE_CANCEL
	CASE_CLOSED = core.CASE_CLOSED
	CASE_STOP   = core.CASE_STOP
)
