package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func say(msg string, rep int) {
	for i := 0; i < rep; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	say("Sync", 10)

	// goroutine is asynchronously and concurrently executed
	go say("Async1", 50)
	go say("Async2", 30)
	go say("Async3", 70)

	time.Sleep(time.Second * 3)

	// using 4 CPUs
	runtime.GOMAXPROCS(4)

	// goroutine using anonymous function & sync.WaitGroup
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("Async-A")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Async-B")

	// wait until wait.Done() is called twice
	// note that wait.Add(2) is stated previously
	wait.Wait()
}
