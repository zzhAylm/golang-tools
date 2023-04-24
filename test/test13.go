package main

import (
	"errors"
	"fmt"
)

func main() {

	// golang 的内置函数：new(int):指针类型
	num1 := new(int)
	*num1 = 100
	fmt.Println(num1, &num1, *num1)

	//	 捕获异常，并且继续往下执行，虽然函数test1报错了，但是可以处理，并且test1后面会继续执行
	test1()

	fmt.Println("-----------")

	err := readConfig("zzh")
	if err != nil {
		//fmt.Println(err)
		// 抛出异常，输出异常，程序终止
		panic(err)
	}
	fmt.Println("-----------")
}
func test1() {
	defer func() {
		// 内置函数，捕获异常，
		if err := recover(); err != nil {
			fmt.Printf("error=%s\n", err)
		}
	}()
	num2 := 0
	num3 := 10
	res := num3 / num2
	fmt.Println(res)

}

func readConfig(name string) (err error) {
	if name == "init.config" {
		return nil
	}
	// 自定义的异常
	return errors.New("读取文件错误。。。。")
}
