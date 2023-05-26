package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Print(s.area(), " ")
	}
	fmt.Println()
}

func showPerimeter(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Print(s.perimeter(), " ")
	}
	fmt.Println()
}

// empty interface can be considered as object type in other language
func printValue(val interface{}) {
	fmt.Println(val)
}

func main() {
	myRect := Rect{15., 30.}
	myCircle := Circle{10.}
	showArea(myRect, myCircle)
	showPerimeter(myRect, myCircle)

	printValue(1)
	printValue("sample string")

	var x interface{} = 12
	y := x
	z := x.(int)
	printValue(y)
	printValue(z)
}
