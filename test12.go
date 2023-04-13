package main

import (
	"fmt"
	"time"
)

// 时间和日期相关函数
func main() {

	now := time.Now()
	fmt.Printf("%v ,%T \n", now, now)

	year := now.Year()
	month := now.Month()
	minute := now.Minute()
	fmt.Println(year, int(month), minute)

	// 格式化：
	sprintf := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	println(sprintf)

	println(now.Format("2006-01-02 15:04:05"))

	// 时间常量
	duration := time.Minute

	println(duration)

	// 毫秒，纳秒
	println(now.UnixNano(), now.Unix())

	// golang 的内置函数：new(int):指针类型

	num1 := new(int)
	*num1 = 100
	fmt.Println(num1, &num1, *num1)

	num2 := 0

	res := *num1 / num2

	fmt.Println(res)
}
