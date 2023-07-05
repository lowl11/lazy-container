package genc

// Len returns length of container
func Len[T any](container []T) int {
	return len(container)
}

// Cap returns capacity of container
func Cap[T any](container []T) int {
	return cap(container)
}

// Empty returns true if container is empty
func Empty[T any](container []T) bool {
	return Len(container) == 0
}

// Clone returns copy of the current container
func Clone[T any](container []T) []T {
	cloneArray := make([]T, Len(container))
	copy(cloneArray, container)
	return cloneArray
}

// Each runs EachFunc typed function for every item in container
func Each[T any](container []T, eachFunc EachFunc[T]) {
	for _, item := range container {
		eachFunc(item)
	}
}

// Where returns container filtered by given WhereFunc
func Where[T any](container []T, whereFunc WhereFunc[T]) []T {
	newContainer := make([]T, 0, len(container))
	for _, item := range container {
		if whereFunc(item) {
			newContainer = append(newContainer, item)
		}
	}
	return newContainer
}

// Count returns quantity of container filtered by given CountFunc
func Count[T any](container []T, countFunc CountFunc[T]) int {
	var counter int
	for _, item := range container {
		if countFunc(item) {
			counter++
		}
	}
	return counter
}

// Contains returns true if given search value is found
func Contains[T any](container []T, searchValue T) bool {
	for _, item := range container {
		if any(item) == any(searchValue) {
			return true
		}
	}

	return false
}

// ContainsFunc returns true if given search value is found with given condition function
func ContainsFunc[T any](container []T, containsFunc ConditionFunc[T]) bool {
	for _, item := range container {
		if containsFunc(item) {
			return true
		}
	}

	return false
}

// IndexOf returns index of found element, if it is not found returns -1
func IndexOf[T any](container []T, searchValue T) int {
	for index, item := range container {
		if any(searchValue) == any(item) {
			return index
		}
	}

	return -1
}

// IndexOfFunc returns index of found element with given condition function, if it is not found returns -1
func IndexOfFunc[T any](container []T, indexOfFunc ConditionFunc[T]) int {
	for index, item := range container {
		if indexOfFunc(item) {
			return index
		}
	}

	return -1
}

// Remove deletes element in container at given index
func Remove[T any](container []T, removeIndex uint) []T {
	if int(removeIndex) > len(container) {
		return container
	}

	return append(container[:removeIndex], container[removeIndex+1:]...)
}
