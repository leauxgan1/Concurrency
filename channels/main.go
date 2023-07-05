package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func exampleOne() {
	myChannel := make(chan int)

	go func() { myChannel <- 1234 }()

	go func() { myChannel <- 6789 }()

	fmt.Printf("The value retrieved from my channel is: %d\n", <-myChannel)

	fmt.Printf("The value retrieved from my channel is: %d\n", <-myChannel)
}

func doHardWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func exampleTwoSlow() {
	mychannel := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			result := doHardWork()
			mychannel <- result
		}
		close(mychannel)
	}()

	for n := range mychannel {
		fmt.Printf("value of n: %d\n", n)
	}
}

func exampleTwoFast() {
	mychannel := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doHardWork()
				mychannel <- result
			}()
		}
		wg.Wait()
		close(mychannel)
	}()

	for n := range mychannel {
		fmt.Printf("value of n: %d\n", n)
	}
}

func main() {
	exampleTwoFast()

}
