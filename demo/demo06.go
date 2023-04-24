package main

import "fmt"

// 基本数据类型
func main() {

	// 布尔数据类型
	var var1 bool = false
	var var2 bool = true
	println(var1)
	fmt.Printf("%T,%t\n", var2, var2)
	fmt.Printf("%T,%t\n", var1, var1)

	// 数值型：
	//		整数：int ,int8
	//		浮点数：float

	var intVar int = 19
	var floatVar float64 = 3.14

	println(intVar, floatVar)

}
