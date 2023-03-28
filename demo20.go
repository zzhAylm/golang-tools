package main

import "fmt"

// 函数数据类型 ，func()、func(int, int) ： 就是一种数据类型
func main() {
	//f1 : 如果不加 括号，f1就是一个变量
	fmt.Printf("f1的类型=%T\n", f1)

	funVar := f1

	sum := funVar(10, 20)

	fmt.Printf("funVar的类型=%T,sum=%d\n", funVar, sum)

	println(f1, funVar)

	f3 := f2

	println(f3, f2)

	// 匿名函数
	f4 := func() int {
		return 1
	}

	println(f4())

	// 匿名函数直接调用执行，返回结果
	i := func() int {
		return 1
	}()

	println(i)

	// 匿名函数自己调用自己，并且传递参数
	func(a, b int) int {
		sum := a + b
		fmt.Println(sum)
		return sum
	}(1, 2)

	// 函数可以当作参数

	i2 := oper(1, 2, addNums)
	println(i2)
	i3 := oper(1, 2, delNums)
	println(i3)

	// 高阶函数传入匿名函数，回调函数
	i4 := oper(1, 1, func(i int, i2 int) int {
		return i / i2
	})
	println(i4)

}

func f1(a, b int) int {
	return a + b
}

func f2() {

}

// 高阶函数，使用函数作为参数
func f3(f func(int, int)) {

}
func addNums(a, b int) int {
	return a + b
}
func delNums(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(int, int) int) int {

	return fun(a, b)
}
