# reentrantlock

This package provides a reentrant mutex implementation in Go using the `sync` package. A reentrant mutex is a mutual exclusion lock that allows the same goroutine to acquire the lock multiple times without deadlocking.

## Usage

To use this package, first create a new reentrant mutex by calling `NewMutex()`:

```go
l := reentrantlock.NewMutex()
```

You can then acquire the lock using the Lock() method:
```go
l.Lock()
```

When you're done with the critical section, you can release the lock using the Unlock() method:
```go
l.Unlock()
```

If a goroutine attempts to acquire a lock it already holds, the lock will be acquired without blocking. The counter in the mutex struct will be incremented.

You can also try to acquire a lock using the TryLock() method, which returns a boolean indicating whether the lock can be acquired:

```go
if l.TryLock() {
// do something
}
```

## Note
- When calling Unlock, if the calling goroutine does not own the lock, the program will panic.
- You need to call `Unlock()` for each time you called `Lock()`.