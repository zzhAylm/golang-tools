package main

import "fmt"

// 函数
func main() {

	num := add(10, 20)
	sum := addNum(10, 20, 20, 30)

	fmt.Printf("add=%d", num)
	fmt.Printf("sum=%d", sum)
}

func add(a int, b int) int {
	return a + b
}

// 多个参数的函数
func addNum(arg ...int) int {

	sum := 0
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}
