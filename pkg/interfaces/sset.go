package interfaces

type SSet[T any] interface {
	Size() int
	Add(x T)
	Remove(x T)
	Find(x T) (T, bool) // a bit different from the original code. Ideally, this code should return T or nil.
}
