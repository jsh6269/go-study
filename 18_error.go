package main

import (
	"fmt"
	"os"
)

// 'error' is a built-in interface type
// type error interface {
//     Error() string
// }

type divisionByZero struct{}

// divisionByZero implements error interface
func (err divisionByZero) Error() string {
	return "cannot divide by 0"
}

func divide(x float64, y float64) (float64, error) {
	if y == 0 {
		return 0, divisionByZero{}
	}
	return x / y, nil
}

func main() {
	_, err := os.Open("C:\\no_such_file\\sample.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := divide(5, 3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	res, err = divide(2, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
