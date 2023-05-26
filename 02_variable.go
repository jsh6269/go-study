package main

func main() {
	var a int
	a = 10

	var b int = 3
	var c, d, e int = 4, 5, 6

	var f float32 = 3.14

	println(a, b, c, d, e, f)

	// 명시되어 있지 않으면 타입을 추론함
	var x, y = 3, 2.1
	println(x, y)

	const pi1 float32 = 3.14
	const pi2 = 3.14

	const (
		name = "suhan"
		age  = 56
	)
	println(name, age)
}
