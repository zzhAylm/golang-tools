package main

import (
	"fmt"
	"io"
	"os"
)

// go语言特性：继承C语音理念（指针），引入包结构，垃圾回收机制，天然并发

// 切片：动态数组
// 延迟执行：defer
// 包结构，指针，并发

func main() {

	//for key, value := range oldMap {
	//	newMap[key] = value
	//}
	//
	//sum := 0
	//for _, value := range array {
	//	sum += value
	//}

	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	a := []rune("zzhAylm")
	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println(string(a))
	var t interface{}
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T 打印 t 有
	case bool:
		fmt.Printf("boolean %t\n", t) // t 的类型为 bool
	case int:
		fmt.Printf("integer %d\n", t) // t 的类型为 int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t 的类型为 *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t 的类型为 *int
	}
}
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// Contents 以字符串形式返回文件的内容。
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	//defer 延迟执行，用于关闭，如果存在多个defer，多个defer之前
	defer f.Close() // f.Close 将在我们运行时运行完了。

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) //append 稍后讨论。
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // 如果我们返回这里，f 将被关闭。
		}
	}
	return string(result), nil // 如果我们返回这里，f 将被关闭。
}
