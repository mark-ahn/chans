package chans

import (
	"context"
	"fmt"
)

type _Prefix_SomeGetter interface {
	Get() (*Some, error)
}
type _Prefix_SomeSetter interface {
	Set(*Some) error
}
type _Prefix_SomeGetSetter interface {
	_Prefix_SomeGetter
	_Prefix_SomeSetter
}

type _Prefix_SomeGetterOfCh struct {
	ctx  context.Context
	r_ch <-chan *Some
}

func (__ *_Prefix_SomeGetterOfCh) Get() (*Some, error) {
	select {
	case <-__.ctx.Done():
		return nil, fmt.Errorf("timeout")
	case get := <-__.r_ch:
		return get, nil
	}
}

type _Prefix_SomeSetterOfCh struct {
	ctx  context.Context
	w_ch chan<- *Some
}

func (__ *_Prefix_SomeSetterOfCh) Set(v *Some) error {
	select {
	case <-__.ctx.Done():
		return fmt.Errorf("timeout")
	case __.w_ch <- v:
		return nil
	}
}

type _Prefix_SomeGetSetterOfCh struct {
	r *_Prefix_SomeGetterOfCh
	w *_Prefix_SomeSetterOfCh
}

func (__ *_Prefix_SomeGetSetterOfCh) Get() (*Some, error) {
	return __.r.Get()
}
func (__ *_Prefix_SomeGetSetterOfCh) Set(v *Some) error {
	return __.w.Set(v)
}

func New_Prefix_SomeGetSetter(ctx context.Context, r <-chan *Some, w chan<- *Some) _Prefix_SomeGetSetter {
	return &_Prefix_SomeGetSetterOfCh{
		r: &_Prefix_SomeGetterOfCh{ctx: ctx, r_ch: r},
		w: &_Prefix_SomeSetterOfCh{ctx: ctx, w_ch: w},
	}
}

// func IntoSomeGetter(r <-chan *Some) _Prefix_SomeGetter {
// 	return
// }
