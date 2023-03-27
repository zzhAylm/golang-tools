package main

// 匿名变量 ,使用 _ 代替变量
func main() {

	a, b := test()
	println(a, b)
	// 如果函数返回的类型，包含很多信息，但是我们只想要其中几个，就可以使用匿名变量，不接受其数据
	// _ ：代表匿名变量，不会分配内存
	c, _ := test()
	println(c)
}

// 自定义函数
func test() (int, int) {

	return 100, 200
}
