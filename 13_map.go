package main

import "fmt"

func main() {
	var myMap map[int]string
	myMap = make(map[int]string)
	myMap[4] = "apple"
	myMap[234] = "orange"
	fmt.Println(myMap[234])

	val, exists := myMap[5]
	if !exists {
		fmt.Println("no such key")
	} else {
		fmt.Println(val)
	}

	grades := map[string]string{
		"jang":  "A",
		"sara":  "C",
		"junho": "A+",
		"babo":  "F",
	}
	delete(grades, "babo")

	for key, val := range grades {
		fmt.Println(key, val)
	}
}
