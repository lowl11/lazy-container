package anyc

// Container wrapper type for slice of any values with a bunch of methods
type Container []any

// New (constructor) creates new container with given slice of any values as argument
func New(container ...any) Container {
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
func (container Container) Contains(searchValue any) bool {
	return Contains(container, searchValue)
}

// IndexOf returns index of found element, if it is not found returns -1
func (container Container) IndexOf(searchValue any) int {
	return IndexOf(container, searchValue)
}
