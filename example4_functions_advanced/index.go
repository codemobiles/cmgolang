package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := "1", "0"
	fmt.Println(a, b)
	a, b = swap(a, b)
	fmt.Println(a, b)
	c, d := split(17)
	fmt.Println(c, d)
}
