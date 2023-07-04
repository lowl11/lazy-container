package interfaces

import "github.com/lowl11/lazy-container/pkg/containers/stringc"

type IString interface {
	// Len returns length of container
	Len() int
	// Cap returns capacity of container
	Cap() int
	// Empty returns true if container is empty
	Empty() bool

	// Join return a string joined by strings in container with given separator
	Join(separator string) string

	// Each runs EachFunc typed function for every item in container
	Each(stringc.EachFunc)
	// Where returns container filtered by given WhereFunc
	Where(stringc.WhereFunc) []string
	// Count returns quantity of container filtered by given CountFunc
	Count(stringc.CountFunc) int

	// Contains returns true if given search value is found
	Contains(string) bool
	// IndexOf returns index of found element, if it is not found returns -1
	IndexOf(string) int

	// Add appends new string value to container
	Add(string)
	// Remove deletes element in container at given index
	Remove(int)

	// Clone returns copy of the current container
	Clone() []string
	// ThreadSafe turns on thread safe mode.
	// Thread safe mode is concurrency
	ThreadSafe()
	// Slice returns slice of strings
	Slice() []string
}
