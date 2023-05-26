package main

import "fmt"

func main() {
	var name string
	var num = 2

	switch num {
	case 1:
		name = "a type"
	case 2:
		name = "b type"
	default:
		name = "other"
	}
	fmt.Println(name)

	switch x := num * 2; x + 1 {
	case 5:
		fmt.Println("expression")
	default:
		fmt.Println("no expression")
	}

	switch x := num * 2; x + 1 {
	case 5:
		fmt.Println("expression")
		fallthrough
	default:
		fmt.Println("arrived here")
	}

	switch {
	case num > 1:
		println("bigger than 1")
	case num > 2:
		println("bigger than 2")
	}

	var value interface{} = 4
	switch value.(type) {
	case int:
		println("int")
	}
}
