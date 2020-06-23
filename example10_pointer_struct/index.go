package main

import "fmt"

type Account struct {
	username, password string
}

func main() {
	acc := Account{"admin", "1234"}
	fmt.Println(acc)
	fmt.Println("Username: ", acc.username, "Password; ", acc.password)

	accPointer := &acc
	fmt.Println(acc, *accPointer)
	fmt.Println(acc, &accPointer.username)
	fmt.Println(acc, &accPointer.username)
	fmt.Println(acc, &accPointer.password)
}
