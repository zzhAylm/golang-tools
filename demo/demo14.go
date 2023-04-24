package main

import "fmt"

func main() {

	// 值传递 ：arr := [4]int{1, 2, 3, 4} ，固定长度的数组

	//引用传递

	// 切片，可以扩容的数组，引用类型, 通过函数修改之后，会改变原来的值
	slice := []int{0, 1, 2, 3}
	fmt.Printf("slice: %v slice addr %p \n", slice, &slice)

	ret := changeSlice(slice)
	fmt.Printf("slice: %v slice addr %p \n", slice, &slice)
	fmt.Printf("slice: %v ret: %v slice addr %p \n", slice, &slice, ret)

	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	updateArray(arr)
	fmt.Println(arr)

	errors := []int{10, 20, 30}
	fmt.Println(errors)
	updatearrays(errors)
	fmt.Println(errors)

}

func changeSlice(s []int) []int {
	s[1] = 111
	return s
}

func updateArray(array []int) {
	array[0] = 10
}

func updatearrays(arrars []int) {
	arrars[0] = 10
}
