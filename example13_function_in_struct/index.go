package main

import "fmt"

type Account struct {
	username, password string
}

func (acc *Account) clear() {
	acc.username = ""
	acc.password = ""
}

func (acc Account) print() {
	fmt.Println(acc)
}

func main() {
	a := Account{"admin", "1234"}
	fmt.Println(a)

	a.clear()
	fmt.Println(a)

	a.print()
}
