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
}
