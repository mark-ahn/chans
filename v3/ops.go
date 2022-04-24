package chans

func DropAll[T any](ch <-chan T) int {
	l := len(ch)
	return Drop(ch, l)
}
func Drop[T any](ch <-chan T, l int) int {
	for i := 0; i < l; i += 1 {
		<-ch
	}
	return l
}

func ConsumeAll[T any](ch <-chan T, f func(index int, value T) error) (int, error) {
	l := len(ch)
	return Consume(ch, f, l)
}

func Consume[T any](ch <-chan T, f func(index int, value T) error, l int) (int, error) {
	var i = 0
	var err error
	for ; i < l; i += 1 {
		err = f(i, <-ch)
		if err != nil {
			return i, err
		}
	}
	return i, nil
}

func PutFull[T any](ch chan<- T, f func(index int) T) int {
	l := cap(ch) - len(ch)
	return Put(ch, f, l)
}
func Put[T any](ch chan<- T, f func(index int) T, l int) int {
	for i := 0; i < l; i += 1 {
		ch <- f(i)
	}
	return l
}

func NewTokken(l int) chan struct{} {
	res := make(chan struct{}, l)
	PutFull(res, func(index int) struct{} { return struct{}{} })
	return res
}
