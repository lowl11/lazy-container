package array_list

import (
	"github.com/lowl11/lazy-container/pkg/interfaces"
	"sync"
)

type Container[T any] struct {
	inner []T

	threadSafe bool
	mutex      sync.Mutex
}

func New[T any](elements ...T) interfaces.IArrayList[T] {
	return &Container[T]{
		inner: elements,
	}
}
