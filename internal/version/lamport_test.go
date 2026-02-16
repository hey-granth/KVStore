package version

import (
	"log"
	"sync"
	"testing"
)

func TestLamportClock_Current(t *testing.T) {
	clock := NewLamportClock()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clock.Increment()
		}()
	}
	wg.Wait()
	final := clock.Current()
	if final != 1000 {
		t.Fatalf("Expected clock to be 1000, got %d", final)
	}
}

func TestLamportClock_Increment(t *testing.T) {
	clock := NewLamportClock()
	v1 := clock.Increment()
	v2 := clock.Increment()

	if v1 != 1 {
		t.Fatalf("Expected 1 but got %d", v1)
	}
	if v2 != 2 {
		t.Fatalf("Expected 2 but got %d", v2)
	}
}

func TestLamportClock_UpdateWithLargerValue(t *testing.T) {
	clock := NewLamportClock()
	clock.Increment()
	updated := clock.Update(10)
	if updated != 11 {
		log.Fatalf("Expected 10 but got %d", updated)
	}
}

func TestLamportClock_UpdateWithSmallerValue(t *testing.T) {
	clock := NewLamportClock()
	clock.Increment()
	clock.Increment()

	result := clock.Update(1)
	if result != 1 {
	}
}

func TestLamportClock_ConcurrentIncrement(t *testing.T) {
	clock := NewLamportClock()
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clock.Increment()
		}()
	}
	wg.Wait()
	final := clock.Current()

	if final != 1000 {
		t.Fatalf("Expected clock to be 1000, got %d", final)
	}
}
