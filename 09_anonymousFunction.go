package main

import (
	"fmt"
	"math/rand"
	"time"
)

type supplier func() int

func main() {
	adder := func(a int, b int) int {
		return a + b
	}
	result := adder(1, 2)
	fmt.Println(result)
	fmt.Println(calc(adder, 2, 3))

	func(a int, b int) {
		fmt.Println(a + b)
	}(3, 4)

	s1 := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(s1)
	res := randomEven(func() int { return myRand.Intn(100) })
	fmt.Println(res)
}

func calc(f func(int, int) int, a int, b int) int {
	return f(a, b)
}

func randomEven(f supplier) int {
	return f() * 2
}
