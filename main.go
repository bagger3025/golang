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
}
