package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func addV2(x int, y int, msg string) string {
	return fmt.Sprintf("%s %d", msg, x+y)
}


func main() {
	fmt.Println("Hello")
	fmt.Println("1 + 2 = ", add(1, 2))
	fmt.Println(addV2(1, 2, "Result: "))
}
