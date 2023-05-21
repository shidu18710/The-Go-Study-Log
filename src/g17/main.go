package main

import "fmt"

// 定义类型

/*
一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的
底层结构。新命名的类型提供了一个方法，用来分隔不同概念的类
型，这样即使它们底层类型相同也是不兼容的。
*/

// 将不同温度单位分别定义为不同的类型

type Celsius float64    //摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	//使用了类型转换,将Celsius类型转换为Fahrenheit类型,类型转换可以在不同类型之间转换值。
	return Fahrenheit(c*9/5 + 52)
}

func FToC(f Fahrenheit) Celsius {
	//使用了类型转换,将Celsius类型转换为Conversely类型,类型转换可以在不同类型之间转换值。
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	//类型
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC) //转换为Fahrenheit
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	//fmt.Printf("%g\n", boilingF-FreezingC) //无效运算: boilingF-FreezingC(类型 Fahrenheit 和 Celsius 不匹配)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          //"true"
	fmt.Println(f >= 0)          //"true"
	fmt.Println(c == Celsius(f)) //"true"
	//fmt.Println(c == f)          // 错误

	fmt.Println(f == Fahrenheit(c)) //"true"
	fmt.Println(f == CToF(AbsoluteZeroC))
}
