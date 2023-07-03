package string_container

import "github.com/lowl11/lazy-container/pkg/containers/stringc"

func (stringContainer *Container) ThreadSafe() {
	stringContainer.threadSafe = true
}

func (stringContainer *Container) Add(value string) {
	stringContainer.lock()
	defer stringContainer.unlock()

	stringContainer.container = append(stringContainer.container, value)
}

func (stringContainer *Container) Remove(index int) {
	stringContainer.lock()
	defer stringContainer.unlock()

	stringContainer.container = stringc.Remove(stringContainer.container, uint(index))
}

func (stringContainer *Container) Len() int {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Len()
}

func (stringContainer *Container) Cap() int {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Cap()
}

func (stringContainer *Container) Empty() bool {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Empty()
}

func (stringContainer *Container) Join(separator string) string {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Join(separator)
}

func (stringContainer *Container) Each(eachFunc stringc.EachFunc) {
	stringContainer.lock()
	defer stringContainer.unlock()

	stringContainer.container.Each(eachFunc)
}

func (stringContainer *Container) Where(whereFunc stringc.WhereFunc) []string {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Where(whereFunc)
}

func (stringContainer *Container) Count(countFunc stringc.CountFunc) int {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Count(countFunc)
}

func (stringContainer *Container) Contains(searchValue string) bool {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Contains(searchValue)
}

func (stringContainer *Container) IndexOf(searchValue string) int {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.IndexOf(searchValue)
}

func (stringContainer *Container) Clone() []string {
	stringContainer.lock()
	defer stringContainer.unlock()

	return stringContainer.container.Clone()
}

func (stringContainer *Container) Slice() []string {
	return stringContainer.container
}
