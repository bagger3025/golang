package main

import "fmt"

func superAdd2(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := superAdd(2, 3, 4, 5, 6, 7)
	fmt.Println(result)
	result = superAdd2(2, 3, 4, 5, 6, 7)
	fmt.Println(result)
}
