package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Zero"
	a[1] = "One"
	fmt.Println(a)
	fmt.Println("1. ", a[0], " 2. ", a[1])

	course1 := [4]string{"Angular", "React", "VueJS", "Flutter"}
	fmt.Println(course1)

	course2 := []string{"Angular", "React", "VueJS", "Flutter", "Android"}
	fmt.Println(course2)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	i := 6
	var s []int = primes[3:i]
	fmt.Println(s)

}
