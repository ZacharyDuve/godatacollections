package godatacollections

type avlTree[K, T any] struct {
	root      avlTreeNode[T]
	tToKFunc  func(T) K
	kCompFunc func(K, K) int
}

type avlTreeNode[T any] struct {
	t     T
	left  *avlTreeNode[T]
	right *avlTreeNode[T]
}

func (this *avlTree[K, T]) Insert(T) error {

}
