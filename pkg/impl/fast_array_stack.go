package impl

import "github.com/marendtx/ods/pkg/interfaces"

type FastArrayStack[T any] struct {
	n int
	a []T
}

func NewFastArrayStack[T any]() interfaces.List[T] {
	return &FastArrayStack[T]{a: make([]T, 0)}
}

func (fas FastArrayStack[T]) Size() int {
	return fas.n
}

func (fas *FastArrayStack[T]) Get(i int) T {
	return fas.a[i]
}

func (fas *FastArrayStack[T]) Set(i int, x T) T {
	y := fas.a[i]
	fas.a[i] = x
	return y
}

func (fas *FastArrayStack[T]) Add(i int, x T) {
	if fas.n+1 > len(fas.a) {
		fas.resize()
	}

	for j := fas.n; j > i; j-- {
		fas.a[j] = fas.a[j-1]
	}
	fas.a[i] = x
	fas.n++
}

func (fas *FastArrayStack[T]) Remove(i int) T {
	x := fas.a[i]
	copy(fas.a[i:fas.n-1], fas.a[i+1:fas.n])
	fas.n--
	if len(fas.a) >= 3*fas.n {
		fas.resize()
	}
	return x
}

func (fas *FastArrayStack[T]) resize() {
	b := make([]T, max(2*fas.n, 1))
	copy(b, fas.a[:fas.n])
	fas.a = b
}
