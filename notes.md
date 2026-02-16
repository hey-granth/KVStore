# Notes for personal reference

## questions i had while building this project

### why are go maps not thread safe?
A map is a hash table, and hash tables are not thread safe because they can be modified by multiple threads at the same time, which can lead to race conditions and data corruption. In Go, maps are not thread safe because they do not have built-in synchronization mechanisms to prevent concurrent access. If multiple goroutines access a map concurrently without proper synchronization, it can lead to unpredictable behavior and potential crashes. To make a map thread safe, you can use a mutex or a sync.Map, which provides concurrent access to the map without the need for explicit locking.

mutex is mutual exclusion, where we reserve a specified resource for a single thread to access at a time. This is done to prevent race conditions and ensure data integrity when multiple threads are accessing shared resources. A mutex allows only one thread to access the resource at a time, while other threads must wait until the mutex is released before they can access the resource. This helps to avoid conflicts and ensures that the shared resource is accessed in a controlled manner.

### what are goroutines?
Goroutines are lightweight threads of execution in the Go programming language. They are managed by the Go runtime and are designed to be efficient and easy to use. Goroutines allow you to run multiple functions concurrently, which can help improve the performance of your program by taking advantage of multiple CPU cores.

### types of mutex in go

- `sync.Mutex`: basic mutex, only one goroutine can hold it at a time. You can use the Lock() and Unlock() methods to acquire and release the mutex.
- `sync.RWMutex`: read-write mutex, allows multiple readers OR one writer. You can use the RLock() and RUnlock() methods for reading and the Lock() and Unlock() methods for writing.
- `sync.WaitGroup`: not a mutex, but a synchronization primitive that allows you to wait for a collection of goroutines to finish. You can use the Add() method to specify the number of goroutines to wait for, and the Done() method to signal that a goroutine has finished.
- `sync.Cond`: condition variable, allows goroutines to wait for certain conditions to be met. You can use the Wait() method to block a goroutine until a condition is met, and the Signal() or Broadcast() methods to wake up waiting goroutines when the condition is met.
- `sync.Once`: ensures that a function is only executed once, even if called from multiple goroutines. You can use the Do() method to execute the function, and it will only run once, regardless of how many times it is called.
- `sync.Pool`: a pool of temporary objects that can be reused to reduce the number of allocations. It is not a mutex, but it can help improve performance by reducing the overhead of creating and destroying objects. You can use the Get() method to retrieve an object from the pool and the Put() method to return an object to the pool for reuse.

### why does `-race` tests so slower than the normal ones?
The `-race` flag in Go enables the race detector, which is a tool that helps identify race conditions in your code. When you run tests with the `-race` flag, the race detector instruments your code to track access to shared variables and detect potential race conditions. This instrumentation adds overhead to the execution of your tests, which is why they run slower than normal tests.