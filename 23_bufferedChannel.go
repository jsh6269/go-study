package main

import "fmt"

func main() {
	// c := make(chan int)
	// c <- 1  // wait until receiver is ready, and it waits forever. deadlock!
	// fmt.Println(<-c)

	ch := make(chan int, 1)
	ch <- 1004
	fmt.Println(<-ch)

	ch = make(chan int, 1)
	sendChan(ch)
	receiveChan(ch)
	sendAndReceiveChan(ch)

	ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// cannot send but can still receive from closed channel
	close(ch)

	// 1st method to get all the values via channel
	// for {
	//     if i, success := <-ch; success {
	//         fmt.Print(i, " ")
	//     } else {
	//         fmt.Println()
	//         break
	//     }
	// }

	// 2nd method to get all the values via channel
	for i := range ch {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func sendChan(ch chan<- int) {
	ch <- 109
}

func receiveChan(ch <-chan int) {
	fmt.Println(<-ch)
}
func sendAndReceiveChan(ch chan int) {
	ch <- 15
	fmt.Println(<-ch)
}
