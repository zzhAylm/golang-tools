package main

import (
	"fmt"
	"strings"
)

// AddNumFunc 闭包 应用实例：类加器：返回的是一个匿名函数，并且匿名函数运用到了一个函数外的的参数，因此这个匿名函数和操作的参数构成闭包
func AddNumFunc() func(x int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

// 闭包 :  makeSuffix 一个类，参数（在类中定意义的变量）是一个属性，返回的匿名函数操作属性
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		name += suffix
		return name
	}

}
func main() {

	numFunc := AddNumFunc()
	fmt.Println(numFunc(1))
	fmt.Println(numFunc(2))
	fmt.Println(numFunc(3))

	// 闭包结构会保留我们传入的 .jpg 结构（类会保存属性的值）
	suffix := makeSuffix(".jpg")
	fmt.Println(suffix("zzh"))
	fmt.Println(suffix("ylm.jpg"))
	fmt.Println(suffix("zzhAylm.JPG"))
}
