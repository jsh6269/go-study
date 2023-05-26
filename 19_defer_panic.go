package main

import (
	"fmt"
	"os"
)

func sayBye(msg string) {
	fmt.Println("Bye: ", msg)
}

func makeFileError() {
	defer sayBye("B")
	f, err := os.Open("no_such_file.txt")
	if err != nil {
		// panic executes all the defer functions which are stated previously (B & A)
		panic(err)
	}
	defer sayBye("C")
	defer f.Close()
}

func testDefer() {
	defer sayBye("X")
	fmt.Println("testDefer()")
}

func main() {
	testDefer()
	fmt.Println("-------------------")
	defer sayBye("A")
	makeFileError()
	defer sayBye("D")

}
