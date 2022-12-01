package main

import (
	"fmt"
	"log"

	"github.com/bagger3025/golang/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("err was", err, "balance is", account.Balance())
}
