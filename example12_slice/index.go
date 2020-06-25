package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[3:4]
	fmt.Println(s)

	slicePrimes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	printSlice(slicePrimes)

	printSlice(slicePrimes[2:])      // Leading slice
	printSlice(slicePrimes[:5])      // Trailing slice
	printSlice(slicePrimes[3 : 3+2]) // Trailing slice

	slicePrimes = append(slicePrimes, 23)
	printSlice(slicePrimes)

	slicePrimes = append(slicePrimes, []int{29, 31}...)
	printSlice(slicePrimes)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
