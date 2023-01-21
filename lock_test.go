package reentrantlock

import (
	"math/rand"
	"sync"
	"testing"
)

func Test_Reentrant(t *testing.T) {

	m := NewMutex()

	l := rand.Intn(30)
	for i := 0; i < l; i++ {
		m.Lock()
	}

	go func() {
		m.Lock()
		t.Error("invalid status")
	}()

	l = rand.Intn(30)
	for i := 0; i < l; i++ {
		m.Lock()
	}
}

func TestMutex_UnlockPanic(t *testing.T) {
	s := &sync.Mutex{}
	s.Lock()

	assertPanic(t, func() {
		m := NewMutex()
		go func() {
			m.Lock()
			s.Unlock()
		}()
		s.Lock()
		m.Unlock()
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestMutex_Unlock(t *testing.T) {
	wg1 := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	m := NewMutex()
	if !m.TryLock() {
		t.Error("invalid status TryLock should return true")
	}

	m.Lock()
	if !m.TryLock() {
		t.Error("invalid status TryLock should return true")
	}

	l := rand.Intn(100)
	for i := 0; i < l; i++ {
		wg1.Add(1)
		wg2.Add(1)
		go func() {
			if m.TryLock() {
				t.Error("invalid status TryLock should return false")
			}
			wg2.Done()
			m.Lock()
			if !m.TryLock() {
				t.Error("invalid status TryLock should return true")
			}
			m.Unlock()
			wg1.Done()
		}()
	}

	wg2.Wait()
	m.Unlock()
	wg1.Wait()
}

func Test_ReentrantRecursive(t *testing.T) {
	m := NewMutex()
	recursive(m, 0, rand.Intn(100), t)
}

func recursive(lock Mutex, counter int, stop int, t *testing.T) {
	lock.Lock()
	if !lock.TryLock() {
		t.Error("invalid status TryLock should return true")
	}

	counter++

	if stop == counter {
		return
	}

	recursive(lock, counter, stop, t)
}
