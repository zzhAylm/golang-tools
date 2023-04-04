package main

import (
	"fmt"
	"main/utils"
)

// 包的介绍：
// 函数的访问修饰符，使用函数名的首字母是否大写来决定的
func main() {

	i := utils.Add(10, 20)
	//i := utils.add(10, 20)
	fmt.Println(i)
}
