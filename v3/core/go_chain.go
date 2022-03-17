package core

import (
	"fmt"
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
var ErrStopIter = ErrStopMap
