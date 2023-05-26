package main

import "fmt"

func main() {
	k := 2
	if k == 1 {
		fmt.Println("one")
	} else if k == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("else")
	}

	if k != 3 && k*2 >= 4 {
		fmt.Println("is true")
	}
	// if문 안에서만 사용가능한 변수를 선언할 수 있음
	if val := k * 2; val > 3 {
		fmt.Println(val, "is bigger than 3")
	}
}
