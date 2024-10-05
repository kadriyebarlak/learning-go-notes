package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//race condition:
//A condition that occurs when multiple goroutines access shared data concurrently,
//and at least one of them modifies the data, leading to unexpected behavior
func main() {
	//raceExample()
	//raceExampleFixed()
	raceExampleFixedWithAtomic()
}

func raceExampleFixedWithAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	var val int32 = 0

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			atomic.AddInt32(&val, 1)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)
}

func raceExampleFixed() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	lock := sync.Mutex{}

	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			lock.Lock()
			val++
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)
}

func raceExample() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	val := 0

	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000000; i++ {
			val++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(val)
}