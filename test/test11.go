package main

import (
	"fmt"
	"strconv"
	"strings"
)

// defer : 延迟、
// 参数传递，值传递和引用传递 ，

// go 中常用的系统函数
func main() {

	var1 := "zzhAylm张紫韩"

	fmt.Println(len(var1))

	// rune类型的切片
	run := []rune(var1)

	// 默认是按自己遍历，转成切片类型就会按字符进行遍历
	for i := 0; i < len(run); i++ {
		fmt.Printf("%c\n", run[i])
	}

	atoi, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误")
	} else {
		fmt.Println(atoi)
	}

	itoa := strconv.Itoa(123)
	fmt.Println(itoa)

	s := "zzhAylm"

	bytes := []byte(s)
	fmt.Println(bytes)
	s2 := string([]byte{97, 98, 99})
	fmt.Println(s2)

	// 十进制转其他进制
	formatInt := strconv.FormatInt(123, 2)
	fmt.Println(formatInt)

	contains := strings.Contains("zzhAylm", "zzh")

	fmt.Println(contains)

	count := strings.Count("zzhAylm", "zzh")
	fmt.Println(count)

	fold := strings.EqualFold("zzhAylm", "ZZHAYLM")
	fmt.Println(fold)

	indexAny := strings.IndexAny("zzhAylm", "zzh")
	fmt.Println(indexAny)

}
