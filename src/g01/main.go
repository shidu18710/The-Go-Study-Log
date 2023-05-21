package main

import (
	"fmt"
	"os"
)

func main() {
	// 定义两个空字符串变量s和sep
	var s, sep string
	//遍历 os.Args 的每个参数 从下标 1 开始
	//获取参数 从下标一开始 [/tmp/go-build2736902484/b001/exe/main a b c d] 懂得
	for i := 1; i < len(os.Args); i++ {
		// 将sep和参数值拼接到s上
		// sep初始为空,然后变成" "用于在参数之间添加空格
		s += sep + os.Args[i]
		sep = " "
	}

	// 打印最终拼接得到的字符串s
	fmt.Println(s)

	fmt.Println(os.Args) //输出：[/tmp/go-build2736902484/b001/exe/main a b c d]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:]) //输出：[a b c d]
}
