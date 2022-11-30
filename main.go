package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func multiply2(a, b int) int {
	return a * b
}

func addAndMultiply(a, b int) (added int, multiplied int) {
	defer fmt.Println(added, multiplied, "are returned")
	added = a + b
	multiplied = a * b
	defer fmt.Println(added, multiplied, "are returned")
	return
}

func repeatme(names ...string) {
	fmt.Println(names)
}

func main() {
	const name = "nico"
	var myName = "hello"
	myName2 := "hello2"

	fmt.Println(name)
	fmt.Println(myName)
	fmt.Println(myName2)

	added, multiplied := addAndMultiply(3, 5)
	fmt.Println(added, multiplied)

	repeatme("a", "b", "c", "d", "e")
}
