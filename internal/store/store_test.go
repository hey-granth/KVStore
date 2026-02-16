package store

import (
	"sync"
	"testing"
)

func TestGetAndSet(t *testing.T) {
	s := NewStore()
	key := "testKey"
	value := Value{Data: "testValue", Version: 1, NodeID: "node1"}

	s.Set(key, value)

	retrieved, exists := s.Get(key)
	if !exists {
		t.Fatalf("Expected key %s to exist", key)
	}
	if retrieved != value {
		t.Fatalf("Expected value %v, got %v", value, retrieved)
	}
}

func AcceptNewerVersion(t *testing.T) {
	s := NewStore()
	key := "testKey"
	oldValue := Value{Data: "oldValue", Version: 1, NodeID: "node1"}
	newValue := Value{Data: "newValue", Version: 2, NodeID: "node2"}

	s.Set(key, oldValue)
	s.Set(key, newValue)

	retrieved, _ := s.Get(key)

	if retrieved != newValue {
		t.Fatalf("Expected value %v, got %v", newValue, retrieved)
	}
}

func RejectOlderVersion(t *testing.T) {
	s := NewStore()
	key := "testKey"
	oldValue := Value{Data: "oldValue", Version: 1, NodeID: "node1"}
	newValue := Value{Data: "newValue", Version: 2, NodeID: "node2"}

	s.Set(key, newValue)
	s.Set(key, oldValue)

	retrieved, _ := s.Get(key)
	if retrieved != newValue {
		t.Fatalf("Expected value %v, got %v", newValue, retrieved)
	}
}

func TestEqualVersionTieBreaker(t *testing.T) {
	s := NewStore()
	key := "testKey"
	value1 := Value{Data: "value1", Version: 1, NodeID: "node1"}
	value2 := Value{Data: "value2", Version: 1, NodeID: "node2"}

	s.Set(key, value1)
	s.Set(key, value2)

	retrieved, _ := s.Get(key)
	if retrieved != value2 {
		t.Fatalf("Expected value %v, got %v", value2, retrieved)
	}
}

func TestConcurrentWrites(t *testing.T) {
	s := NewStore()
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key"
			value := Value{Data: "value", Version: int64(i), NodeID: "node"}
			s.Set(key, value)
		}(i)
	}
	wg.Wait()

	_, exists := s.Get("key")
	if !exists {
		t.Fatalf("Expected key to exist after concurrent writes")
	}
}
