package version

import "sync"

type LamportClock struct {
	mu    sync.Mutex
	clock int64
}

func NewLamportClock() *LamportClock {
	return &LamportClock{}
}

func (l *LamportClock) Increment() int64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.clock++
	return l.clock
}

func (l *LamportClock) Update(received int64) int64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	if received > l.clock {
		l.clock = received
	}
	l.clock++
	return l.clock
}

func (l *LamportClock) Current() int64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.clock
}

// lamport clocks are never meant to be decreased, so we don't need a method to decrement the clock. The Update method will ensure that the clock is always updated to the maximum of the current clock and the received clock, and then incremented by one. This ensures that the Lamport clock always moves forward in time, even when receiving messages from other nodes.
