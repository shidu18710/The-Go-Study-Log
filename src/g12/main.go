package main

import "fmt"

func main() {
	p := new(int)
	*p = 1
	fmt.Println(*p)
}
