package main

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}

func main() {
	//变量逃逸
	f()
	g()
	println(global)
	println(*global)
}
