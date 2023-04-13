package main

import (
	"fmt"
	"main/utils"
)

// 包的介绍：
// 函数的访问修饰符，使用函数名的首字母是否大写来决定的

// 构建项目： go build -o ～/my.exe xxx/project/main

// 匿名函数，函数也是一中数据类型

var (
	addFunc = func(a, b int) int {
		return a + b
	}
)

func main() {

	i := utils.Add(10, 20)
	//i := utils.add(10, 20)
	fmt.Println(i)

	// 匿名函数
	funcVar := func(a, b int) int {
		return a + b
	}(11, 20)
	fmt.Println(funcVar)

	// 将函数赋给一个变量
	funcAdd := func(a, b int) int {
		return a + b
	}

	i2 := funcAdd(10, 20)
	fmt.Println(i2)

	// 全局匿名函数
	i3 := addFunc(12, 7)
	fmt.Println(i3)
}
