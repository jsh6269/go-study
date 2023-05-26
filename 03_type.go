package main

import "fmt"

func main() {
	var a bool = true
	var b1 string = "hello"
	var b2 string = "Hello, " + "World!"
	var b3 string = `hello\nworld`
	var b4 string = "hello\nworld"
	fmt.Println(a)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println("-------------")
	fmt.Println(b4)
	fmt.Println("-------------")

	// 타입 생략해도 byte로 해석. 즉 65가 출력됨
	var b5 byte = 'A'
	fmt.Println(b5)

	var x int = 4
	var xx int8 = 4
	var xxx int16 = 4
	var xxxx int32 = 4
	var xxxxx int64 = 4
	var y uint = 4
	var yy uint8 = 4
	var yyy uint16 = 4
	var yyyy uint32 = 4
	var yyyyy uint64 = 4
	fmt.Println(x, xx, xxx, xxxx, xxxxx)
	fmt.Println(y, yy, yyy, yyyy, yyyyy)

	var f1 float32 = 3.14
	var f2 float64 = 3.14
	fmt.Println(f1, f2)

	var m int = int(f1)
	str2 := []byte(b1)
	fmt.Println(m)
	fmt.Println(str2)
}
