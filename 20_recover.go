package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("no_such_file")
	println("Done")
}

func openFile(file string) {
	// defer is executed when panic is called
	defer func() {
		// panic is expired due to recover function call
		if r := recover(); r != nil {
			fmt.Println("<error description>", r)
		}
	}()

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
