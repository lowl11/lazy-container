package array_list

import "github.com/lowl11/lazy-container/pkg/interfaces"

func (container *Container[T]) lock() {
	if !container.threadSafe {
		return
	}

	container.mutex.Lock()
}

func (container *Container[T]) unlock() {
	if !container.threadSafe {
		return
	}

	container.mutex.Unlock()
}

func (container *Container[T]) copy() interfaces.IArrayList[T] {
	newContainer := New[T]()
	
	if container.threadSafe {
		newContainer.ThreadSafe()
	}

	return newContainer
}
