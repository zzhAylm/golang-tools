package main

import "fmt"

// 类型转换，都是显示的转换
func main() {

	a := 5
	b := 5.0

	println(a, b)

	c := float64(a)
	println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
