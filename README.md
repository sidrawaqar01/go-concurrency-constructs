# go-concurrency-constructs

The project contains exmaples of below constructs in order: 

**Go Routine:** A lightweight thread of execution in Go. It allows concurrent operations, enabling functions or methods to run independently of each other. Goroutines are managed by the Go runtime and are easy to create using the go keyword.

**Wait Group:** A synchronization primitive used to wait for a collection of goroutines to finish executing. It provides methods like Add, Done, and Wait to manage and block until the goroutines complete.

**Mutex:** Short for "mutual exclusion," it's a locking mechanism used to ensure that only one goroutine can access a particular section of code at a time. This prevents data races when multiple goroutines access shared resources.

**RWMutex (Read-Write Mutex):** A special type of mutex that allows multiple readers to access a resource simultaneously, but only one writer can access it at a time. It provides methods like RLock (read lock) and Lock (write lock), enabling better performance in read-heavy scenarios.

**Cond (Condition Variable):** A synchronization primitive used for signaling between goroutines. It allows one or more goroutines to wait until a condition is met. The Wait and Signal methods are typically used to coordinate tasks between goroutines.

**Once:** A type that ensures a function is only executed once, no matter how many times it is called. It is useful for setting up initialization tasks that must be done exactly once, such as in a singleton pattern.

**Pool:** A collection of pre-allocated resources (often goroutines, memory, or buffers) that can be reused, reducing the overhead of creating and destroying objects frequently. The sync.Pool type is used for managing object reuse in Go.

**Channels:** A communication mechanism that allows goroutines to exchange data safely. Channels can be used to send and receive values between goroutines, allowing synchronization and coordination between them.

**GoMaxProcs (GOMAXPROCS):** A runtime variable that determines the maximum number of CPU cores that Go can utilize for executing goroutines concurrently. By default, it's set to the number of CPUs, but it can be adjusted to control concurrency.
