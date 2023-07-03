package stringc

func Len(container []string) int {
	return len(container)
}

func Cap(container []string) int {
	return cap(container)
}

func Empty(container []string) bool {
	return Len(container) == 0
}

func Clone(container []string) []string {
	cloneArray := make([]string, Len(container))
	for i := 0; i < len(container); i++ {
		cloneArray[i] = container[i]
	}
	return cloneArray
}

func Each(container []string, eachFunc EachFunc) {
	for _, item := range container {
		eachFunc(item)
	}
}

func Where(container []string, whereFunc WhereFunc) []string {
	newContainer := make([]string, 0, len(container))
	for _, item := range container {
		if whereFunc(item) {
			newContainer = append(newContainer, item)
		}
	}
	return newContainer
}

func Count(container []string, countFunc CountFunc) int {
	var counter int
	for _, item := range container {
		if countFunc(item) {
			counter++
		}
	}
	return counter
}

func Contains(container []string, searchValue string) bool {
	for _, item := range container {
		if searchValue == item {
			return true
		}
	}

	return false
}

func IndexOf(container []string, searchValue string) int {
	for index, item := range container {
		if searchValue == item {
			return index
		}
	}

	return -1
}

func Remove(container []string, removeIndex uint) []string {
	if int(removeIndex) > len(container) {
		return container
	}

	return append(container[:removeIndex], container[removeIndex+1:]...)
}
