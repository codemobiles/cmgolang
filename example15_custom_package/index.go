package main

import (
	test "example15/add"
	"fmt"
	color "github.com/fatih/color"
)

func main() {

	result := test.Add(5, 6)
	msg := fmt.Sprintf("Result : %d", result)
	color.Red(msg)
}
