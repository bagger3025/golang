package main

import "fmt"

func main() {
	numbers := [5]int{2, 3, 4, 5, 6}

	numbers2 := []int{2, 3}

	numbers2 = append(numbers2, 4)

	fmt.Println(numbers)

	fmt.Println(numbers2)
}
