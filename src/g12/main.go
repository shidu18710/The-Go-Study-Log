package main

import "fmt"

func main() {
	//new创建初始变量地址
	p := new(int)
	*p = 1
	fmt.Println(p)
	fmt.Println(*p)
}
