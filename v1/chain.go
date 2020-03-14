package chans

import (
	"github.com/mark-ahn/chans/v1/core"
)

var ErrSkipMap = core.ErrSkipMap
var ErrStopMap = core.ErrStopMap
var ErrStopIter = core.ErrStopIter

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
