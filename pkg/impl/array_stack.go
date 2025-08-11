package impl

import "github.com/marendtx/ods/pkg/interfaces"

type ArrayStack[T any] struct {
	n int
	a []T
}

func NewArrayStack[T any]() interfaces.List[T] {
	return &ArrayStack[T]{a: make([]T, 0)}
}

func (as ArrayStack[T]) Size() int {
	return as.n
}

func (as *ArrayStack[T]) Get(i int) T {
	return as.a[i]
}

func (as *ArrayStack[T]) Set(i int, x T) T {
	y := as.a[i]
	as.a[i] = x
	return y
}

func (as *ArrayStack[T]) Add(i int, x T) {
	if as.n+1 > len(as.a) {
		as.resize()
	}

	for j := as.n; j > i; j-- {
		as.a[j] = as.a[j-1]
	}
	as.a[i] = x
	as.n++
}

func (as *ArrayStack[T]) Remove(i int) T {
	x := as.a[i]
	for j := i; j < as.n-1; j++ {
		as.a[j] = as.a[j+1]
	}
	as.n--
	if len(as.a) >= 3*as.n {
		as.resize()
	}
	return x
}

func (as *ArrayStack[T]) resize() {
	b := make([]T, max(2*as.n, 1))
	for i := range as.n {
		b[i] = as.a[i]
	}
	as.a = b
}
