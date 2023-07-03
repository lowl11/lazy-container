package string_container

func (stringContainer *Container) lock() {
	if !stringContainer.threadSafe {
		return
	}

	stringContainer.mutex.Lock()
}

func (stringContainer *Container) unlock() {
	if !stringContainer.threadSafe {
		return
	}

	stringContainer.mutex.Unlock()
}
