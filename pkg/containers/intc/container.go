package intc

type Container []int

func New(container ...int) Container {
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

func (container Container) Clone() Container {
	return Clone(container)
}

func (container Container) Sum() int {
	return Sum(container)
}
