package stringc

// Len returns length of container
func Len(container []string) int {
	return len(container)
}

// Cap returns capacity of container
func Cap(container []string) int {
	return cap(container)
}

// Empty returns true if container is empty
func Empty(container []string) bool {
	return Len(container) == 0
}

// Clone returns copy of the current container
func Clone(container []string) []string {
	cloneArray := make([]string, Len(container))
	copy(cloneArray, container)
	return cloneArray
}

// Each runs EachFunc typed function for every item in container
func Each(container []string, eachFunc EachFunc) {
	for _, item := range container {
		eachFunc(item)
	}
}

// Where returns container filtered by given WhereFunc
func Where(container []string, whereFunc WhereFunc) []string {
	newContainer := make([]string, 0, len(container))
	for _, item := range container {
		if whereFunc(item) {
			newContainer = append(newContainer, item)
		}
	}
	return newContainer
}

// Count returns quantity of container filtered by given CountFunc
func Count(container []string, countFunc CountFunc) int {
	var counter int
	for _, item := range container {
		if countFunc(item) {
			counter++
		}
	}
	return counter
}

// Contains returns true if given search value is found
func Contains(container []string, searchValue string) bool {
	for _, item := range container {
		if searchValue == item {
			return true
		}
	}

	return false
}

// IndexOf returns index of found element, if it is not found returns -1
func IndexOf(container []string, searchValue string) int {
	for index, item := range container {
		if searchValue == item {
			return index
		}
	}

	return -1
}

// Remove deletes element in container at given index
func Remove(container []string, removeIndex uint) []string {
	if int(removeIndex) > len(container) {
		return container
	}

	return append(container[:removeIndex], container[removeIndex+1:]...)
}
