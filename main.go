package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	nico := person{"nico", 18, []string{"kimchi", "anythingelse"}}
	nico2 := person{
		name:    "nico",
		age:     18,
		favFood: []string{"kimchi", "bulgogi"},
	}
	fmt.Println(nico)
	fmt.Println(nico.favFood)

	fmt.Println(nico2)
	fmt.Println(nico2.favFood)
}
