package array_list

import (
	"github.com/lowl11/lazy-container/pkg/containers/genc"
	"github.com/lowl11/lazy-container/pkg/interfaces"
	"sort"
)

func (container *Container[T]) ThreadSafe() interfaces.IArrayList[T] {
	container.threadSafe = true
	return container
}

func (container *Container[T]) Len() int {
	container.lock()
	defer container.unlock()

	return genc.Len[T](container.inner)
}

func (container *Container[T]) Cap() int {
	container.lock()
	defer container.unlock()

	return genc.Cap[T](container.inner)
}

func (container *Container[T]) Empty() bool {
	container.lock()
	defer container.unlock()

	return genc.Empty[T](container.inner)
}

func (container *Container[T]) Push(values ...T) interfaces.IArrayList[T] {
	container.lock()
	defer container.unlock()

	container.inner = append(container.inner, values...)
	return container
}

func (container *Container[T]) Remove(index int) interfaces.IArrayList[T] {
	container.lock()
	defer container.unlock()

	container.inner = genc.Remove[T](container.inner, uint(index))
	return container
}

func (container *Container[T]) Get(getFunc genc.GetFunc[T]) *T {
	index := container.IndexOf(func(item T) bool {
		return getFunc(item)
	})

	if index < 0 {
		return nil
	}

	return &container.inner[index]
}

func (container *Container[T]) Contains(conditionFunc genc.ConditionFunc[T]) bool {
	container.lock()
	defer container.unlock()

	return genc.ContainsFunc[T](container.inner, conditionFunc)
}

func (container *Container[T]) IndexOf(conditionFunc genc.ConditionFunc[T]) int {
	container.lock()
	defer container.unlock()

	return genc.IndexOfFunc[T](container.inner, conditionFunc)
}

func (container *Container[T]) Where(whereFunc genc.WhereFunc[T]) interfaces.IArrayList[T] {
	container.lock()
	defer container.unlock()

	return container.copy().Push(genc.Where[T](container.inner, whereFunc)...)
}

func (container *Container[T]) Each(eachFunc genc.EachFunc[T]) interfaces.IArrayList[T] {
	container.lock()
	defer container.unlock()

	genc.Each[T](container.inner, eachFunc)
	return container
}

func (container *Container[T]) Count(countFunc genc.CountFunc[T]) int {
	container.lock()
	defer container.unlock()

	return genc.Count[T](container.inner, countFunc)
}

func (container *Container[T]) Sort(sortFunc genc.SortFunc[T]) interfaces.IArrayList[T] {
	container.lock()
	defer container.unlock()

	sort.Slice(container.inner, sortFunc)
	return container
}

func (container *Container[T]) Clone() interfaces.IArrayList[T] {
	container.lock()
	defer container.unlock()

	return container.copy().Push(container.inner...)
}
