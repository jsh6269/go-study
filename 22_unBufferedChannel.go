package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello, World!"
	}()

	str := <-ch
	fmt.Println(str)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
		done <- true
	}()
	// receiver waits until value is passed through the channel
	<-done
	fmt.Println("Process ended")
}
