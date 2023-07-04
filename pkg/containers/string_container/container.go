package string_container

import (
	"github.com/lowl11/lazy-container/pkg/containers/stringc"
	"github.com/lowl11/lazy-container/pkg/interfaces"
	"sync"
)

// Container is a string container using stringc.Container
type Container struct {
	container stringc.Container

	threadSafe bool
	mutex      sync.Mutex
}

// New (constructor) creates new instance of Container with given slice of strings as argument
func New(values ...string) interfaces.IString {
	return &Container{
		container: stringc.New(values...),
	}
}
