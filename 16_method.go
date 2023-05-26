package main

import "fmt"

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r Rect) changeWidth_useless(newWidth int) {
	r.width = newWidth
}

func (r *Rect) changeWidth_useful(newWidth int) {
	r.width = newWidth
}

func main() {
	myRect := Rect{15, 12}
	fmt.Println(myRect.area())

	// value receiver cannot change actual field value
	myRect.changeWidth_useless(5)
	fmt.Println(myRect.width)

	// pointer receiver cannot change actual field value
	myRect.changeWidth_useful(5)
	fmt.Println(myRect.width)
}
