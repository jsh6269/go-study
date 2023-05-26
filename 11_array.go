package main

import "fmt"

func main() {
	var a [5]string
	a[0] = "hi"
	a[1] = "my"
	a[2] = "is"
	a[3] = "su"
	a[4] = "han"

	for i := 0; i < 5; i++ {
		fmt.Print(a[i] + " ")
	}
	fmt.Println()

	var b = [2]int{1, 2}
	var c = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b[1])
	fmt.Println(c[3])

	var multArr [2][3][4]int
	multArr[1][2][3] = 5

	var x = [2][3]int{
		{0, 1, 2},
		// when the input is broken into several lines, golang tries to add semicolon at the end of the line
		// to prevent that, add comma eventhough it is the last element
		{3, 4, 5},
	}
	fmt.Println(x[0][1])
}
