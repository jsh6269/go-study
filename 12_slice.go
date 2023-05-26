package main

import "fmt"

func main() {
	// slice
	var a []int
	a = []int{1, 2, 3, 4}
	a[1] = 5
	fmt.Println(a)

	// length = 6, capacity = 10
	s := make([]int, 6, 10)
	fmt.Println(s)

	// slice s.t. len = cap = 0 is considered as 'nil'
	var n []int
	fmt.Println(n == nil)

	// sub-slice
	fmt.Println(a[1:3])
	fmt.Println(a[2:])
	fmt.Println(a[:2])

	// append values
	fmt.Println(a)
	a = append(a, 3)
	a = append(a, 4, 5, 6, 7)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	// append slice (merge)
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6, 7}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	// copy
	copy(slice1, slice2)
	fmt.Println(slice1)

	slice2_copy := make([]int, len(slice2), cap(slice2))
	copy(slice2_copy, slice2)
	fmt.Println(slice2_copy)

}
