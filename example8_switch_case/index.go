package main

import "fmt"

func main() {

	var lim int

	for true {
		fmt.Print("Please enter limit : ")
		fmt.Scanf("%d", &lim)
		switch lim {
		case 0:
			fmt.Println("Limit is zero")
		case 1:
			fmt.Println("Limit is one")
		case 2:
			fmt.Println("Limit is two")
		default:
			fmt.Println("Limit is somethingelse")
			return
		}
	}

	fmt.Println("Program exit!")
}
