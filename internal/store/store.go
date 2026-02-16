package store

import "sync"

type Value struct {
	Data    string // the actual value stored
	Version int64  // logical clock version, used for conflict resolution
	NodeID  string // used as tie-breaker when versions are equal
}

type Store struct {
	mu   sync.RWMutex     // RWMutex allows multiple concurrent read operations OR a single write operation at a time, ensuring thread safety for the in-memory store
	data map[string]Value // actual in memory storage
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]Value),
	}
}

// Get : getting a value from the store, returns the value and a boolean indicating if the key exists
func (s *Store) Get(key string) (Value, bool) {
	s.mu.RLock()         // acquire read lock for concurrent reads
	defer s.mu.RUnlock() // ensure the lock is released after reading
	value, exists := s.data[key]
	return value, exists
}

func (s *Store) Set(key string, incoming Value) {
	s.mu.Lock()         // acquire write lock for exclusive access
	defer s.mu.Unlock() // ensure the lock is released after writing
	existing, exists := s.data[key]

	if !exists {
		s.data[key] = incoming
		return
	}
	if shouldReplace(existing, incoming) {
		s.data[key] = incoming
	}
}

func shouldReplace(v Value, in Value) bool {
	if in.Version > v.Version {
		return true
	}
	if in.Version == v.Version {
		return in.NodeID > v.NodeID // tie-breaker based on NodeID
	}
	return false
}
