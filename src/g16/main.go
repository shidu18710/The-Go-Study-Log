package main

func main() {
	//可赋值性 隐式的赋值行为和nil

	//创建一个切片时，发生隐式的赋值行为
	medals := []string{"gold", "silver", "bronze"}
	/***********  隐形的赋值  ***********/
	medals[0] = "gold"
	medals[1] = "silver"
	medals[2] = "bronze"

	//不管是隐式还是显式地赋值，在赋值语句左边的变量和右边最终的求 到的值必须有相同的数据类型。

	//nil可以赋值给任何指针或引用类型的变量
	//引用类型有：指针、slice切片、管道channel、接口interface、map、函数等
	a := 1
	p := &a
	p = nil
	println(p)
	medals = nil
}
