package godatacollections

type Stack[T any] interface {
	Push(T)
	Pop() T
	Len() int
}
