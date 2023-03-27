package main

// 变量的作用域： 定义在函数外面的成为全局变量，定义在函数体内的变量为局部变量，
// 全局变量和局部变量是可以重名，在函数体内会有就近原则，优先使用局部变量名

// 常量 ：

// 特殊常量； iota ,自动计数常量
func main() {
	const name string = "zzh"
	const name1 = "zzh2"
	const constVar int = 18

	const var1, var2, var3 = "zzh", 10, "hebei"
	println("常量="+name, name1, constVar)
	println(var1, var2, var3)

	var a1, b2, c3 = "zzh", 19, "hebei"
	println(a1, b2, c3)

	const (
		a = iota
		b
		c
		d
		e
		f = "zzh"
		g
		i = 100
		j
		k = iota
	)
	println(a, b, c, d, e, f, g, i, j, k)
}
