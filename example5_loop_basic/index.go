package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)

		if sum > 500 {
			// return
			break
		}
	}
	fmt.Println(sum)
}
