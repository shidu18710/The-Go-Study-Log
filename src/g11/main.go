package main

import (
	"flag"
	"fmt"
	"strings"
)

// 添加 参数，设置参数名称，参数默认值，参数的作用
var n = flag.Bool("n", false, "不知到啊")
var sep = flag.String("s", " ", "你知道吗")

func main() {
	//flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）
	flag.Parse()
	//flag.Args() 返回命令行参数中不包含标志的部分，并以字符串切片的形式返回。
	//程序中的 sep 和 n 变量分别是指向对应命令行标志参数变量的指针，因此必须用 *sep 和 *n 形式的指针语法间接引用它们
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
