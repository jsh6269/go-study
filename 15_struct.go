package main

import "fmt"

type student struct {
	name  string
	age   int
	grade string
}

type dict struct {
	data map[string]string
}

func newDict() *dict {
	// constructor function
	// map object should be initialized
	d := dict{}
	d.data = make(map[string]string)
	return &d
}

func main() {
	a := student{}
	b := student{"jang", 21, "A"}
	c := student{name: "junho", grade: "A+"}
	d := new(student)
	a.age = 25
	d.name = "sara"

	var students = [...]student{a, b, c, *d}
	for i := 0; i < 4; i++ {
		fmt.Print("s", i, ": ")
		fmt.Println(students[i].name, students[i].age, students[i].grade)
	}

	myDict := newDict()
	myDict.data["jang"] = "A"
	fmt.Println(myDict.data["jang"])
}
