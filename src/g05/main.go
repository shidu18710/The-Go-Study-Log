package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "shiDu")
	} else {
		for _, arg := range files {
			//os.Open() 打开指定的文件
			/*
				第一个值是被打开的文件( *os.File ），其后被 Scanner 读取。
				第二个值是内置 error 类型的值。如果 err 等于内置值 nil （译注：相当于其它语言里的NULL），那么文件被成功打开。
			*/
			f, err := os.Open(arg)
			if err != nil {
				//处理错误
				fmt.Fprintf(os.Stdin, "dup2: %v\n", err)
			}
			countLines(f, counts, arg)
			/*
				调用 Close 关闭该文件，并释放占用的所有资源。
			*/
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// *os.File 表示一个打开的文件
// 在Go语言中,用来表示指针。os.File是一个表示打开文件的结构体类型,os.File指针()表示这个类型的一个具体实例。
// 所以*os.File可以传递给读取该文件内容的函数,让函数可以操作这个打开的文件对象。
func countLines(f *os.File, counts map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for list, n := range counts {
		if n > 1 {
			println("值:", list, "文件名", fileName)
		}
	}
}
