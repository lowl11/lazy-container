package interfaces

import "github.com/lowl11/lazy-container/pkg/containers/genc"

type IArrayList[T any] interface {
	Push(...T) IArrayList[T]
	Remove(int) IArrayList[T]

	Get(genc.GetFunc[T]) *T

	Contains(genc.ConditionFunc[T]) bool
	IndexOf(genc.ConditionFunc[T]) int

	Sort(genc.SortFunc[T]) IArrayList[T]
	Where(genc.WhereFunc[T]) IArrayList[T]
	Each(genc.EachFunc[T]) IArrayList[T]
	Count(genc.CountFunc[T]) int

	Len() int
	Cap() int
	Empty() bool
	ThreadSafe() IArrayList[T]

	Clone() IArrayList[T]
}
