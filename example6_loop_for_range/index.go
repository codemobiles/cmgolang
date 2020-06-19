package main

import "fmt"

func main() {
	data := []string{"a", "b", "c"}

	for i, item := range data {
		fmt.Println(i, item)
	}
}
