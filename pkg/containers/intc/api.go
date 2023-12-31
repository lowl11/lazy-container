package intc

// Len returns length of container
func Len(container []int) int {
	return len(container)
}

// Cap returns capacity of container
func Cap(container []int) int {
	return cap(container)
}

// Empty returns true if container is empty
func Empty(container []int) bool {
	return Len(container) == 0
}

// Clone returns copy of the current container
func Clone(container []int) []int {
	cloneArray := make([]int, Len(container))
	copy(cloneArray, container)
	return cloneArray
}

// Each runs EachFunc typed function for every item in container
func Each(container []int, eachFunc EachFunc) {
	for _, item := range container {
		eachFunc(item)
	}
}

// Where returns container filtered by given WhereFunc
func Where(container []int, whereFunc WhereFunc) []int {
	newContainer := make([]int, 0, len(container))
	for _, item := range container {
		if whereFunc(item) {
			newContainer = append(newContainer, item)
		}
	}
	return newContainer
}

// Count returns quantity of container filtered by given CountFunc
func Count(container []int, countFunc CountFunc) int {
	var counter int
	for _, item := range container {
		if countFunc(item) {
			counter++
		}
	}
	return counter
}

// Contains returns true if given search value is found
func Contains(container []int, searchValue int) bool {
	for _, item := range container {
		if searchValue == item {
			return true
		}
	}

	return false
}

// IndexOf returns index of found element, if it is not found returns -1
func IndexOf(container []int, searchValue int) int {
	for index, item := range container {
		if searchValue == item {
			return index
		}
	}

	return -1
}

// Remove deletes element in container at given index
func Remove(container []int, removeIndex uint) []int {
	if int(removeIndex) > len(container) {
		return container
	}

	return append(container[:removeIndex], container[removeIndex+1:]...)
}

// Sum returns sum of all container values
func Sum(container []int) int {
	var sum int
	for _, item := range container {
		sum += item
	}
	return sum
}
