package reentrantlock

import (
	"sync"
)

type mutex struct {
	owner   int
	counter int
	mutex   *sync.Mutex
}

type Mutex interface {
	Lock()
	Unlock()
	TryLock() bool
}

func NewMutex() Mutex {
	return &mutex{
		mutex: &sync.Mutex{},
		owner: -1,
	}
}

func (m *mutex) Lock() {
	id := goroutineId()
	if id == m.owner {
		m.counter++
		return
	}

	m.mutex.Lock()
	m.counter++
	m.owner = id
}

func (m *mutex) Unlock() {
	if m.owner != goroutineId() {
		panic("unlock of not owned mutex")
	}

	m.counter--
	if m.counter != 0 {
		return
	}

	m.owner = -1
	m.mutex.Unlock()
}

func (m *mutex) TryLock() bool {
	return m.owner == -1 || goroutineId() == m.owner
}
