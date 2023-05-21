package main

import "fmt"

var q = f()

func f() *int {
	v := 1
	return &v
}

func main() {
	var i, j, k int
	var a string = "11"
	fmt.Println(i, j, k, a)

	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var b, c int
	fmt.Println(&b == &b, &c == &b, &b != nil)

	fmt.Println(q == f())
}
