package main

import "fmt"

func main() {
	var a, b = 2, 3
	adder(a, b)
	add5(&a)
	fmt.Println(a)
	printAll("My", "name", "is", "suhan")
	res := addAll1(2, 3, 4, 5)
	fmt.Println()
	fmt.Println(res)
	c, r := addAll2(2, 3, 4, 5)
	fmt.Println(c, r)
	c, r = addAll3(2, 3, 4, 5)
	fmt.Println(c, r)
}

// pass by value
func adder(var1 int, var2 int) {
	fmt.Println(var1 + var2)
}

// pass by reference
func add5(x *int) {
	*x = *x + 5
}

func printAll(msg ...string) {
	for _, s := range msg {
		fmt.Print(s + " ")
	}
}

func addAll1(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func addAll2(nums ...int) (int, int) {
	var cnt int = 0
	var sum int = 0
	for _, num := range nums {
		cnt++
		sum += num
	}
	return cnt, sum
}

func addAll3(nums ...int) (cnt int, sum int) {
	// cnt, sum = 0, 0  (automatically initialized to 0)
	for _, num := range nums {
		cnt++
		sum += num
	}
	return
}
