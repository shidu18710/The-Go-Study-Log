package main

import "fmt"

func main() {
	//认识 new

	// new 创建初始变量地址赋值给p （p指针）
	p := new(int)
	//
	*p = 1
	fmt.Println(p)
	fmt.Println(*p)
}
