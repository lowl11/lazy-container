package stringc

import "strings"

// Container wrapper type for slice of strings with a bunch of methods
type Container []string

// New (constructor) creates new container with given slice of strings as argument
func New(container ...string) Container {
	return container
}

// Len returns length of container
func (container Container) Len() int {
	return Len(container)
}

// Cap returns capacity of container
func (container Container) Cap() int {
	return Cap(container)
}

// Empty returns true if container is empty
func (container Container) Empty() bool {
	return Empty(container)
}

// Join return a string joined by strings in container with given separator
func (container Container) Join(separator string) string {
	return strings.Join(container, separator)
}

// Each runs EachFunc typed function for every item in container
func (container Container) Each(eachFunc EachFunc) {
	Each(container, eachFunc)
}

// Where returns container filtered by given WhereFunc
func (container Container) Where(whereFunc WhereFunc) Container {
	return Where(container, whereFunc)
}

// Count returns quantity of container filtered by given CountFunc
func (container Container) Count(countFunc CountFunc) int {
	return Count(container, countFunc)
}

// Clone returns copy of the current container
func (container Container) Clone() Container {
	return Clone(container)
}

// Contains returns true if given search value is found
func (container Container) Contains(searchValue string) bool {
	return Contains(container, searchValue)
}

// IndexOf returns index of found element, if it is not found returns -1
func (container Container) IndexOf(searchValue string) int {
	return IndexOf(container, searchValue)
}
