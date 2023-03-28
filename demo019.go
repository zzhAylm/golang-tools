package main

import "fmt"

type T struct {
	name  string // 对象的名称
	value int    // 它的值
}

// 函数 defer : 推迟、延迟 ,当有多个defer语句时会逆序执行，一般用于关闭操作
func main() {

	sum := getSum(5)
	fmt.Println(sum)

	f("1")
	f("2")
	f("3")
	defer f("4")
	f("5")
	defer f("6")
	f("7")
	defer f("8")
	f("9")
}

func getSum(num int) int {
	if num <= 1 {
		return num
	}
	return num + getSum(num-1)
}

func f(s string) {
	fmt.Println(s)
}
