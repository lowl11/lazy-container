package intc

func Len(container []int) int {
	return len(container)
}

func Cap(container []int) int {
	return cap(container)
}

func Empty(container []int) bool {
	return Len(container) == 0
}

func Clone(container []int) []int {
	cloneArray := make([]int, Len(container))
	for i := 0; i < len(container); i++ {
		cloneArray[i] = container[i]
	}
	return cloneArray
}

func Sum(container []int) int {
	var sum int
	for _, item := range container {
		sum += item
	}
	return sum
}
