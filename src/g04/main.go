package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		dup 的第一个版本打印标准输入中多次出现的行，以重复次数开头。
		该程序将引入 if 语句， map 数据类型以及 bufio 包。
	*/

	/*
		map存储了键/值（key/value）的集合
		在例子中键是字符串，值是整数
		内置函数 make 创建空 map
	*/
	counts := make(map[string]int)
	/*
		input 是一个 bufio.Scanner,用于从 os.Stdin 读取输入。os.Stdin 是标准输入。
		bufio.Scanner 是一个用于读取输入的结构体类型。它提供了一些方法用于扫描输入,读取行等。
	*/
	input := bufio.NewScanner(os.Stdin)
	//Scan() 读取输入的下一行,并将其保存到Scannern的缓冲区。返回值会告诉你读取是否成功。
	/*
		1. 调用input.Scan()读取标准输入的下一行
		2. 如果读取成功,则执行循环体的语句
		3. 如果读取失败(例如达到文件末尾),则退出循环
	*/
	for input.Scan() {
		//Text():返回以Scan()方法读取的最新行
		//将当前读取的单词(input.Text() 返回的字符串)作为 key  lookup counts map。然后将对应的值加一,实现计数作用。
		//重复的不断加 1，不重复的成为 1。
		counts[input.Text()]++
		/*
			line := input.Text()
			counts[line] = counts[line] + 1
		*/
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
