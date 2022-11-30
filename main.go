package main

import "fmt"

func canIDrink(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIDrink2(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	default:
		return true
	}
	return false
}

func canIDrink3(age int) bool {
	switch koreanage := age + 2; {
	case koreanage < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	default:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(20))
	fmt.Println(canIDrink(10))
}
