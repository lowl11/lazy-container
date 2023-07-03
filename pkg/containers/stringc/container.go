package stringc

import "strings"

type Container []string

func New(container ...string) Container {
	return container
}

func (container Container) Len() int {
	return Len(container)
}

func (container Container) Cap() int {
	return Cap(container)
}

func (container Container) Empty() bool {
	return Empty(container)
}

func (container Container) Join(separator string) string {
	return strings.Join(container, separator)
}

func (container Container) Each(eachFunc EachFunc) {
	Each(container, eachFunc)
}

func (container Container) Where(whereFunc WhereFunc) []string {
	return Where(container, whereFunc)
}

func (container Container) Count(countFunc CountFunc) int {
	return Count(container, countFunc)
}

func (container Container) Clone() Container {
	return Clone(container)
}

func (container Container) Contains(item string) bool {
	return Contains(container, item)
}

func (container Container) IndexOf(searchValue string) int {
	return IndexOf(container, searchValue)
}
