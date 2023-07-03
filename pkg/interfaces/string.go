package interfaces

import "lazy-container/pkg/containers/stringc"

type IString interface {
	Len() int
	Cap() int
	Empty() bool

	Join(separator string) string

	Each(stringc.EachFunc)
	Where(stringc.WhereFunc) []string
	Count(stringc.CountFunc) int

	Contains(string) bool
	IndexOf(string) int

	Add(string)
	Remove(int)

	Clone() []string
	ThreadSafe()
	Slice() []string
}
