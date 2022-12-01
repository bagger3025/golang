package main

import (
	"fmt"

	"github.com/bagger3025/golang/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
