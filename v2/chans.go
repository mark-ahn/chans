package chans

import (
	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in broadcaster__template.go -out broadcaster__template__gen.go gen "_Prefix_=Of Some=Bytes,BUILTINS,interface{},struct{}"
type Some generic.Type
type Other generic.Type

type Bytes = []byte

type DoneNotifier interface {
	Done() <-chan struct{}
}
