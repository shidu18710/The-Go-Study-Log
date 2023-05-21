package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 定义两个空字符串变量s和sep
	s, sep := "", ""

	// 遍历os.Args[1:] 获取命令行参数(除第一个之外)
	/*
		1，_ 的作用
		2，每次循环迭代， range 产生一对值；索引以及在该索引处的元素值。
		3，但 range 的语法要求, 要处理元素, 必须处理索引
		4，一种思路是把索引赋值给一个临时变量, 如 temp , 然后忽略它的
		值，但Go语言不允许使用无用的局部变量
		5，Go语言中这种情况的解决方法是用 空标识符 （blank identifier），
		即 _ （也就是下划线）。
		6，空标识符可用于任何语法需要变量名但程序
		逻辑不需要的时候, 例如, 在循环里，丢弃不需要的循环索引, 保留元
		素值。
	*/
	for _, arg := range os.Args[1:] {
		// 将sep和参数值拼接到s上
		// sep初始为空,然后变成" "用于在参数之间添加空格
		s += sep + arg
		sep = " "
	}
	// 打印最终拼接得到的字符串s
	fmt.Println(s)

	/*
		它接收一个字符串切片和一个sep分隔符字符串,并将切片的所有元素连接成一个大字符串,在元素之间添加sep分隔符
		- strings 是 os.Args[1:],包含了除第一个参数之外的所有命令行参数
		- sep 是 " ",是一个空格字符串
	*/
	fmt.Println(strings.Join(os.Args[1:], " ")) //输出：a b c d

	for list, arg := range os.Args[1:] {
		fmt.Printf("索引：%v\t值：%v\n", list, arg)
	}
}
