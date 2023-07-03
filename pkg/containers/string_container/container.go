package string_container

import (
	"github.com/lowl11/lazy-container/pkg/containers/stringc"
	"github.com/lowl11/lazy-container/pkg/interfaces"
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
