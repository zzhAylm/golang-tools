package main

import "fmt"

// 编码  ASCII 、GBK、Unicode编码
func main() {

	var var1 = 'a'
	fmt.Printf("%d", var1)

	// 连接符 +
	fmt.Print("hello" + "zzh")
	// 转义字符 \
	fmt.Print("hello \" zzh ")
}
