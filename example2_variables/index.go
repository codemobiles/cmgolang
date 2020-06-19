package main

import "fmt"

var tmp1 string
var tmp2 bool
var tmp3 int = 10
var tmp4 = "go programming"

var v1, v2, v3 string

func main() {
	tmp1 = "codemobiles"
	tmp2 = true
	fmt.Println(tmp1, tmp2)
	tmp5 := 6666

	t1, t2, t3 := 1, 2, 3

	fmt.Println("tmp1 = ", tmp1, tmp2, tmp4, tmp5)
	fmt.Println("tx = ", t1, t2, t3)
	t3 = t3 + 1

	const t4 int = 5
	// t4 = t4 + 1
	fmt.Println("tx = ", t1, t2, t3)

}
