package main

import "fmt"

func main() {
	var a, b int = 7, 5
	fmt.Println(a+b, a-b, a/b, a%b)
	fmt.Println(a == b, a != b, a > b, a <= b)
	fmt.Println(a != b && a > b, a == b || a > b+1)
	fmt.Println(!(a != b))
	fmt.Println(a << 2)
	fmt.Println(a&b, a|b, a^b)

	a += 2
	a++
	fmt.Println(a)

	var x int = 5
	var ptr = &x
	fmt.Println(*ptr)
	fmt.Println(ptr)
}
