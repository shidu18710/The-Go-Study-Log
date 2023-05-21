package main

// 结构体

type Books struct {
	title string
}

func main() {
	//赋值

	//命名变量的赋值
	var x = 1
	println(x)

	//通过指针间接赋值
	var p *bool
	tr := true
	p = &tr
	*p = false
	println(p)
	println(*p)

	//结构体字段赋值
	var person Books
	person.title = "Go 语言"
	println(person.title)

	//数组、slice或map的元素赋值
	var balance [10]float64
	balance[0] = 2
	println(balance[0])
}
