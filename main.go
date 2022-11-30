package main

import "fmt"

func main() {
	nico := map[string]string{"name": "nico", "age": "18"}

	fmt.Println(nico)

	for key, val := range nico {
		fmt.Println(key, val)
	}
}
