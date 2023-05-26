package testlib

import "fmt"

func init(){
	fmt.Println("init function is executed")
}

// function name should start with upper case to make it 'public'
func GetGreeting(){
	fmt.Println("Hello, World!")
}