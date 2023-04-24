package main

import (
	"fmt"
	"unsafe"
)

// golang 字符 ， 字符变量 UTF-8
func main() {
	var z int = '中'

	fmt.Printf("%d,%c", z, z)

	var str = "zzh"
	fmt.Println(str)

	fmt.Printf("%c", str[0])
	fmt.Printf("%c", str[1])
	fmt.Printf("%c", str[2])

	// 字符使用%c 输出 unicode 字符，没有一个字符都是一个整型，但其都对应一个 unicode 字符，我们使用%c进行输出
	//
	// 字符类型和进行运算，字符本质就是一个整型的数

	b := false

	fmt.Println("b所占的空间大小", unsafe.Sizeof(b))

	// 字符串 , 使用反引号，输出文本，不对内容进行编译
	// 双引号，转义字符 \ 转移特殊符号
	// 字符串拼接，需要用 +  如果用欧多行，+需要放在上一行，否则 go 会在每一行的末尾加一个分号，语法就会不合法
	str = `package main

			import "fmt"
			
			// 字符串
			func main() {
			
				var str = "zzhAylm"
			
				fmt.Printf("str长度=%d\n", len(str))
			
				fmt.Printf("str[3]=%d\n", str[3])
				fmt.Printf("str[3]=%c\n", str[3])
			}`
	fmt.Println(str)

	//类型转换： 大的值转为小的值，会造成溢出，内容和我们想象的不一样

	numb1 := 9999
	numb2 := int8(numb1)
	fmt.Println(numb1, numb2)

	// byte 的取值范围是 0-255

	// int8 ： 的取值范围 -128-127
	// int16 ：
	// int32
	// int64

	// uint8 和 byte 一样： 0-255
	// uint16
	// uint32
	// uint64
	var n1 byte = 255
	fmt.Println(n1)

	// 大的数值类型不能直接赋给其他的数据类型

	var n64 int64 = 120
	var n8 int8 = int8(n64) + 1

	fmt.Println(n8)

}
