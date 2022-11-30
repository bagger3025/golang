package main

import "fmt"

func canIDrink(age int) bool {
	if koreanage := age + 2; koreanage < 18 {
		return false
	} else {
		return true
	}
}
func main() {
	fmt.Println(canIDrink(20))
	fmt.Println(canIDrink(10))
}
