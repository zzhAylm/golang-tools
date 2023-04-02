package main

import "fmt"

func main() {
	// 数组：
	var a [10]int
	a[0] = 10
	fmt.Println(a)

	// 切片 , 传入到其他函数中，会改变原来的值，地址传递，切片就像数组的引用
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	updateQpArr(s)
	qpArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	fmt.Println(qpArr)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// a1或b的修改，也会改变names中的值
	// 他们用的同一份地址
	a1 := names[0:2]
	b := names[1:3]
	fmt.Println(a1, b)

	b[0] = "XXX"
	fmt.Println(a1, b)
	fmt.Println(names)
}

func updateQpArr(arr []int) {
	arr[0] = 10

}
