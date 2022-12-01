package main

import (
	"fmt"

	"github.com/bagger3025/golang/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
