package string_container

import (
	"lazy-container/pkg/containers/stringc"
	"lazy-container/pkg/interfaces"
	"sync"
)

type Container struct {
	container stringc.Container

	threadSafe bool
	mutex      sync.Mutex
}

func New(values ...string) interfaces.IString {
	return &Container{
		container: stringc.New(values...),
	}
}
