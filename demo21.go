package main

// 闭包结构

/*
一个外层函数中，有内层函数，并且内层函数操作外层函数的局部变量
并且外层函数，返回的就是这个内层函数，这个内层函数和外层函数的局部变量，统称为闭包结构
局部变量的生命周期会发生变化，正常的局部变量会随着函数的调用而创建，随着函数结束而销毁，
但是闭包结构中的外层函数的局部变量不会随着外层函数的结束而销毁，因为内层函数还在继续使用
*/
func main() {

	f4 := increment()

	println(f4())
	println(f4())
	println(f4())
	println(f4())
	println(f4())

	f5 := increment()
	println(f5())
	println(f5())
	println(f5())
	println(f5())

}

func increment() func() int {
	i := 0
	f4 := func() int {
		i++
		return i
	}
	return f4
}
