package main

func main() {
	//元组赋值

	x := 1
	y := 2

	x, y = y, x
	//输出交换内容
	println(x, y)

	//计算两个整数值的的最大公约数
	println(gcd(1, 4))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
