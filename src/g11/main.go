package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "不知到啊")
var sep = flag.String("s", " ", "你知道吗")

func main() {
	flag.Parse()
	//flag.Args() 返回命令行参数中不包含标志的部分，并以字符串切片的形式返回。
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
