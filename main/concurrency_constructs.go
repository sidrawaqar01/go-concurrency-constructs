package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
1. go routine
2. wait group
3. mutex
4. rwmutex
5. cond
6. once
7. pool
8. channels
9. gomaxprolever
*/

func goRoutineDefinitions() {

	// way 1
	go sampleFuncForGoRoutine()

	// way 2 - it defines and runs the go routine at same time, it can not be reused because its anonymous
	go func() {
		fmt.Println("this is go routine 2")
	}()

	// way 3 - define variable as function then call it as go routine
	sampleFunctionForGoRoutine := func() {
		fmt.Println("this is go routine 3")
	}
	go sampleFunctionForGoRoutine()

}

func sampleFuncForGoRoutine() {
	fmt.Println("this is go routine 1")
}

func waitGroupSimpleExample() {

	var wg sync.WaitGroup // wait group is used to wait for all go routines to finish

	sayHelloWithWaitGroup := func() {
		defer wg.Done()
		fmt.Println("waitGroupSimpleExample ran")
	}

	wg.Add(1) // add 1 to wait group - this tells it has to wait for one go routine to finish
	go sayHelloWithWaitGroup()
	wg.Wait() // wait for all go routines to finish - this will block further execution until all go routines are finished
}

func waitGroupCanModifyVariableOutsideItsScope() {
	var wg sync.WaitGroup
	fruit := "apple"

	wg.Add(1) // we add .Add(1) here because otherwise wg.Wait() will be called before go routine is added to wait group
	go func() {
		defer wg.Done()
		fruit = "banana"
	}()

	wg.Wait()

	fmt.Printf("fruit name is: %s\n", fruit)
}

func waitGroupLoopWithWronglyUsedVariable() {

	var salutationWaitgroupLoop sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		salutationWaitgroupLoop.Add(1)
		go func() {
			defer salutationWaitgroupLoop.Done()
			fmt.Println(salutation)
		}()
	}
	salutationWaitgroupLoop.Wait()

}

func waitGroupLoopWithCorrectlyUserVariable() {

	var salutationWaitgroupLoop sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		salutationWaitgroupLoop.Add(1)
		go func(s string) {
			defer salutationWaitgroupLoop.Done()
			fmt.Println(s)
		}(salutation)
	}
	salutationWaitgroupLoop.Wait()
}

func mutex() {

	var countProtectedByMutex int
	var mutex sync.Mutex
	var waitGroupForMutex sync.WaitGroup

	increment := func() {
		mutex.Lock()
		defer mutex.Unlock()
		countProtectedByMutex++
		fmt.Printf("count is: %d\n", countProtectedByMutex)
	}

	decrement := func() {
		mutex.Lock()
		defer mutex.Unlock()
		countProtectedByMutex--
		fmt.Printf("count is: %d\n", countProtectedByMutex)
	}

	for i := 0; i < 5; i++ {
		waitGroupForMutex.Add(1)
		go func() {
			defer waitGroupForMutex.Done()
			increment()
			decrement()
		}()
	}

	waitGroupForMutex.Wait()
}

func rwmutex() {
	var countProtectedByRWMutex int
	var mutexRW sync.RWMutex
	var waitGroupForMutexRW sync.WaitGroup

	start := time.Now()

	incrementRW := func() {
		mutexRW.Lock()
		defer mutexRW.Unlock()
		time.Sleep(100 * time.Millisecond)
		countProtectedByRWMutex++
		fmt.Printf("count after increment is: %d and time is: %v\n", countProtectedByRWMutex, time.Since(start))
	}

	readCountProtectedByRWMutex := func() {
		mutexRW.RLock()
		defer mutexRW.RUnlock()
		fmt.Printf("count after read is: %d and time is: %v\n", countProtectedByRWMutex, time.Since(start))
	}

	for i := 0; i < 5; i++ {
		waitGroupForMutexRW.Add(1)
		go func() {
			defer waitGroupForMutexRW.Done()
			incrementRW()
		}()
	}

	for i := 0; i < 5; i++ {
		waitGroupForMutexRW.Add(1)
		go func() {
			defer waitGroupForMutexRW.Done()
			readCountProtectedByRWMutex()
		}()
	}
	waitGroupForMutexRW.Wait()
}

func waitingForGoRoutineToFinish() {
	dummyCondition := false

	go func() {
		time.Sleep(2 * time.Second)
		dummyCondition = true
	}()

	for dummyCondition == false {
		fmt.Println("waiting for dummyCondition to be true")
	}

	for dummyCondition == false {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("waiting for dummyCondition to be true with wait")
	}
}

func conditionWithSignal() {

	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	dummyCondition := false

	start := time.Now()

	go func() {
		cond.L.Lock()
		time.Sleep(2 * time.Second)
		dummyCondition = true
		cond.L.Unlock()
		cond.Signal()
	}()

	cond.L.Lock()
	for dummyCondition == false {
		cond.Wait()
	}

	fmt.Println("dummyCondition is true after: ", time.Since(start))

	cond.L.Unlock()
}

// skip this for now
func conditionWithBroadcast() {

	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Broadcast()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func once() {
	once := sync.Once{}
	count := 0

	increment := func() {
		count++
	}

	decrement := func() {
		count--
	}

	for i := 0; i < 100; i++ {
		once.Do(increment)
		once.Do(decrement)
	}
	fmt.Printf("count is %d\n", count)
}

func pool() {

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new instance")
			return struct{}{}
		},
	}

	pool.Get()             // this will return an instance if that exists, otherwise will create one
	instance := pool.Get() // will create one more instance
	pool.Put(instance)     // will put the instance back to the pool
	pool.Get()             // will return the instance that was put back, hence not initializing a new one
}

func unbufferedChannelsTwoWay() {

	twoWayChannel := make(chan string)
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		twoWayChannel <- "hello"
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		fmt.Println(<-twoWayChannel)
	}()

	defer waitGroup.Wait()
}

func unbufferedChannelsOneWay() {

	twoWayChannel := make(chan string)

	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)
	go func(receiver chan<- string) {
		defer waitGroup.Done()
		receiver <- "hello"
	}(twoWayChannel)

	waitGroup.Add(1)
	go func(sender <-chan string) {
		defer waitGroup.Done()
		fmt.Println(<-sender)
	}(twoWayChannel)

	defer waitGroup.Wait()
}

func channelCloseExampleToGiveSignalToBlockedGoroutines() {
	twoWayChannel := make(chan interface{})
	waitGroup := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			<-twoWayChannel // this will block until the channel is closed
			fmt.Println(i)
		}(i)
	}

	// this will close the channel and all the goroutines will be unblocked
	// this is similar to broadcast in condition variable
	close(twoWayChannel)

	waitGroup.Wait()

	//close(twoWayChannel) // this will result in infinite waiting as we are doing close after wait ;(
}

func bufferedChannels() {
	twoWayChannel := make(chan string, 5)
	waitGroup := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println(<-twoWayChannel)
		}()
	}

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			twoWayChannel <- fmt.Sprint("hello ", i)
		}(i)
	}

	waitGroup.Wait()
}

func rangeOverChannel1() {

	channel3 := make(chan string, 5)

	for i := 0; i < 5; i++ {
		channel3 <- fmt.Sprint("hello channel3 ", i)
	}
	close(channel3)

	for c := range channel3 {
		fmt.Println(c)
	}
}

func rangeOverChannel2() {

	channel := make(chan string, 5)

	for _, c := range []string{"hello", "world", "channel"} {
		select {
		case channel <- c:
		}
	}
	close(channel)

	for c := range channel {
		fmt.Println(c)
	}
}

func selectStatement() {

	channel1 := make(chan string, 3)
	channel2 := make(chan string, 3)

	for i := 0; i < 3; i++ {
		channel1 <- fmt.Sprint("hello from channel1 ", i)
		channel2 <- fmt.Sprint("hello from channel2 ", i)
	}

	for i := 0; i < 10; i++ {
		select {
		case value := <-channel1:
			fmt.Println(value)
		case value := <-channel2:
			fmt.Println(value)
		case <-time.After(10 * time.Millisecond):
			fmt.Println("Timed out.")
		default:
			fmt.Println("no message received")
		}
	}
}

func setRuntimeGOMAXPROCS() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
