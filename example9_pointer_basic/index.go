package main

import "fmt"

func main() {
	i := 42
	p := &i
	var p2 *int
	p2 = p
	p2 = &i

	fmt.Println("I = ", *p)
	fmt.Println("I = ", *p2)
	fmt.Println("Addres iof I = ", p)

	*p = *p2 / 2
	fmt.Println("I = ", *p)

	mutate(p)
	fmt.Println("I = ", *p)

	mutate(&i)
	fmt.Println("I = ", *p)
}

func mutate(_p *int) {
	*_p = *_p + 1
}
