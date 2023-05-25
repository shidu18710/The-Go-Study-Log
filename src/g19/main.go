package main

import "go02/src/g19/popcount"

//包的初始化首先是解决包级变量的依赖顺序，然后安照包级变量声明
//出现的顺序依次初始化：

var a = b + c // a 第三个初始化, 为 3
var b = f()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
var c = 1     // c 第一个初始化, 为 1

func f() int {
	return c + 1
}

func main() {
	//包的初始化
	println(a)
	println(b)
	println(c)
	popcount.PopCount(12)
}
