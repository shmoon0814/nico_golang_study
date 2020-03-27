package main

import (
	"fmt"
	"github.com/shmoon0814/learngo/bank_account/accounts"
)

// Account struct
func main() {
	account := accounts.NewAccount("상혁")
	fmt.Println(account)

	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.WithDraw(30)

	if err != nil {
		fmt.Println(err)
	}
	//error handler를 이렇게 작성을 직접 해줘야 한다...

	fmt.Println(account.Balance())
}
