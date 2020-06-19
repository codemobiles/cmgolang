package main

import "fmt"

func main() {

	amount := 10
	if amount > 0 {
		fmt.Println("Amount > 0")
	} else {
		fmt.Println("Amount <= 0")
	}

	if max := 5; max > amount {
		fmt.Println("max > Amount")
	} else {
		fmt.Println("max =< Amount")
	}

	a := "care1"
	b := "lek"

	if a == "care" || b != "lek" {
		fmt.Println("a is equal care")
	} else {
		fmt.Println("a is not equal care")
	}

	// && and ||

}
