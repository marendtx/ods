package interfaces

type List[T any] interface {
	Size() int
	Get(i int) T
	Set(i int, x T)
	Add(i int, x T)
	Remove(i int)
}
