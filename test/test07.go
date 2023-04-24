package main

import (
	"fmt"
	"strconv"
)

// 数据类型的转换
func main() {
	n1 := 100
	// 其他数据类型转为str
	sprintf := fmt.Sprintf("int64转为str%d\n", n1)
	fmt.Printf("sprintf的类型:%T ， str=%s", sprintf, sprintf)

	n2 := 100.0
	fmt.Printf("n2的类型:%T ， n2=%f\n", n2, n2)

	str := fmt.Sprintf("float转为str%f\n", n2)
	fmt.Printf("str的类型:%T ， str=%s\n", str, str)

	// %q: 将字符串加上双引号
	fmt.Printf("str的类型:%T ， str=%q\n", str, str)

	formatInt := strconv.FormatInt(100, 10)
	fmt.Printf("str的类型:%T ， str=%q\n", formatInt, formatInt)

	// str to 基本数据类型的注意事项： 如果失败了会直接返回默认值，而不是报错

	// *int 指针类型
	//  var i int= 100 i指向一个变量  &i是这个变量的地址
	// var p *int = &i  p 执行一个地址，地址的值是 &i  &i的值是100   *p 获取他指向 的地址值的指向的值

	//
}
