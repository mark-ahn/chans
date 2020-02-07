// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package chans

import "context"

func OfFuncBool(ctx context.Context, f func() bool, n int) <-chan bool {
	ch := make(chan bool, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncBoolSlice(ctx context.Context, f func() []bool, n int) <-chan []bool {
	ch := make(chan []bool, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncByte(ctx context.Context, f func() byte, n int) <-chan byte {
	ch := make(chan byte, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncByteSlice(ctx context.Context, f func() []byte, n int) <-chan []byte {
	ch := make(chan []byte, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncComplex128(ctx context.Context, f func() complex128, n int) <-chan complex128 {
	ch := make(chan complex128, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncComplex128Slice(ctx context.Context, f func() []complex128, n int) <-chan []complex128 {
	ch := make(chan []complex128, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncComplex64(ctx context.Context, f func() complex64, n int) <-chan complex64 {
	ch := make(chan complex64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncComplex64Slice(ctx context.Context, f func() []complex64, n int) <-chan []complex64 {
	ch := make(chan []complex64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncError(ctx context.Context, f func() error, n int) <-chan error {
	ch := make(chan error, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncErrorSlice(ctx context.Context, f func() []error, n int) <-chan []error {
	ch := make(chan []error, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncFloat32(ctx context.Context, f func() float32, n int) <-chan float32 {
	ch := make(chan float32, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncFloat32Slice(ctx context.Context, f func() []float32, n int) <-chan []float32 {
	ch := make(chan []float32, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncFloat64(ctx context.Context, f func() float64, n int) <-chan float64 {
	ch := make(chan float64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncFloat64Slice(ctx context.Context, f func() []float64, n int) <-chan []float64 {
	ch := make(chan []float64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt(ctx context.Context, f func() int, n int) <-chan int {
	ch := make(chan int, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncIntSlice(ctx context.Context, f func() []int, n int) <-chan []int {
	ch := make(chan []int, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt16(ctx context.Context, f func() int16, n int) <-chan int16 {
	ch := make(chan int16, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt16Slice(ctx context.Context, f func() []int16, n int) <-chan []int16 {
	ch := make(chan []int16, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt32(ctx context.Context, f func() int32, n int) <-chan int32 {
	ch := make(chan int32, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt32Slice(ctx context.Context, f func() []int32, n int) <-chan []int32 {
	ch := make(chan []int32, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt64(ctx context.Context, f func() int64, n int) <-chan int64 {
	ch := make(chan int64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt64Slice(ctx context.Context, f func() []int64, n int) <-chan []int64 {
	ch := make(chan []int64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt8(ctx context.Context, f func() int8, n int) <-chan int8 {
	ch := make(chan int8, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInt8Slice(ctx context.Context, f func() []int8, n int) <-chan []int8 {
	ch := make(chan []int8, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncRune(ctx context.Context, f func() rune, n int) <-chan rune {
	ch := make(chan rune, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncRuneSlice(ctx context.Context, f func() []rune, n int) <-chan []rune {
	ch := make(chan []rune, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncString(ctx context.Context, f func() string, n int) <-chan string {
	ch := make(chan string, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncStringSlice(ctx context.Context, f func() []string, n int) <-chan []string {
	ch := make(chan []string, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint(ctx context.Context, f func() uint, n int) <-chan uint {
	ch := make(chan uint, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUintSlice(ctx context.Context, f func() []uint, n int) <-chan []uint {
	ch := make(chan []uint, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint16(ctx context.Context, f func() uint16, n int) <-chan uint16 {
	ch := make(chan uint16, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint16Slice(ctx context.Context, f func() []uint16, n int) <-chan []uint16 {
	ch := make(chan []uint16, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint32(ctx context.Context, f func() uint32, n int) <-chan uint32 {
	ch := make(chan uint32, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint32Slice(ctx context.Context, f func() []uint32, n int) <-chan []uint32 {
	ch := make(chan []uint32, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint64(ctx context.Context, f func() uint64, n int) <-chan uint64 {
	ch := make(chan uint64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint64Slice(ctx context.Context, f func() []uint64, n int) <-chan []uint64 {
	ch := make(chan []uint64, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint8(ctx context.Context, f func() uint8, n int) <-chan uint8 {
	ch := make(chan uint8, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUint8Slice(ctx context.Context, f func() []uint8, n int) <-chan []uint8 {
	ch := make(chan []uint8, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUintptr(ctx context.Context, f func() uintptr, n int) <-chan uintptr {
	ch := make(chan uintptr, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncUintptrSlice(ctx context.Context, f func() []uintptr, n int) <-chan []uintptr {
	ch := make(chan []uintptr, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInterface(ctx context.Context, f func() interface{}, n int) <-chan interface{} {
	ch := make(chan interface{}, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncInterfaceSlice(ctx context.Context, f func() []interface{}, n int) <-chan []interface{} {
	ch := make(chan []interface{}, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncStruct(ctx context.Context, f func() struct{}, n int) <-chan struct{} {
	ch := make(chan struct{}, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}

func OfFuncStructSlice(ctx context.Context, f func() []struct{}, n int) <-chan []struct{} {
	ch := make(chan []struct{}, n)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
			}

			t := f()

			select {
			case <-ctx.Done():
				break loop
			default:
			}

			ch <- t
		}
	}()
	return ch
}
