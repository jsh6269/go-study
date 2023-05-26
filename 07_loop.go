package main

import "fmt"

func main() {
	// ordinary for loop
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	// for loop looks like while
	sum = 0
	var n int = 1
	for n <= 100 {
		sum += n
		n++
	}
	fmt.Println(sum)

	// continue, break
	n, sum = 1, 0
	for {
		if n%2 == 0 {
			n++
			continue
		}
		sum += n
		if sum >= 625 {
			break
		}
		n++
	}
	fmt.Println(n)

	// goto
	goto END
	fmt.Println("skipped")
END:
	fmt.Println("end")

}
