package main

import "fmt"

// 变量的定义 ,var 变量名 类型=值 ，
// 变量赋值 变量名 :=值 ，自动识别类型
func main() {

	var test string = "zzh"
	var str = "zzhStr"
	var a, b, c int
	var (
		name    string
		age     int
		address string
	)

	var1 := "zzh"
	var2 := 18

	var3 := "var3"

	println(test)
	println(str)
	println(name, age, address)

	println(a, b, c)

	// 打印类型
	fmt.Printf("%T,%T\n ", var1, var2)

	fmt.Printf("name=%s,age=%d\n", var1, var2)

	// &变量名 ： 获取变量的地址
	fmt.Printf("var=%s,内存地址=%p\n ", var3, &var3)
}
