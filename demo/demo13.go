package main

import "fmt"

// 参数传递类型，引用数据类型，值类型
func main() {
	//值传递 :  整型数组

	arr := [4]int{1, 2, 3, 4}

	fmt.Println(arr)

	updateArr(arr)

	fmt.Println(arr)

}

func updateArr(arrArg [4]int) {
	arrArg[0] = 10
}
