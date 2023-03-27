package main

// 变量交换
func main() {
	var1 := 100
	var2 := 200

	println(var1, var2)
	var1, var2 = var2, var1

	println(var1, var2)

}
