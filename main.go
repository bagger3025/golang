package main

import (
	"fmt"

	mydicts "github.com/bagger3025/golang/mydict"
)

func main() {
	dictionaries := mydicts.Dictionary{"first": "First word"}
	val, err := dictionaries.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	err2 := dictionaries.Add("hello", "world")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(dictionaries)
	}

	err3 := dictionaries.Update("hello", "Golang!")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(dictionaries["hello"])
	}

	err4 := dictionaries.Delete("hello")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("hello is deleted")
	}
}
