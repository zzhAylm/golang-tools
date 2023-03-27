package main

import "fmt"

// 运算符 ，位运算符： &   ｜  ^  >>  <<
func main() {

	var1 := 60
	var2 := 13

	var3 := var1 & var2

	var4 := var1 | var2

	var5 := var1 ^ var2

	var5 += var1
	println(var3)
	println(var4)

	fmt.Printf("%d,二进制=%b\n", var3, var3)
	fmt.Printf("%d,二进制=%b\n", var4, var4)
	fmt.Printf("%d,二进制=%b\n", var5, var5)

	var var6 int
	var var7 int

	fmt.Scanln(&var6, &var7)
	fmt.Printf("var6=%d,var7=%d", var6, var7)
}
