package main

import "fmt"

// for  while goto
func main() {

	str := "zzhAylm!张紫韩"

	// 传统的遍历方式，是按照字节进行遍历的，但是我们的汉字 是三个字节，只取一个字节无法遍历汉字
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d,%c\n", i, str[i])
	}

	// 传统字节的改进
	slice := []rune(str)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d,%c\n", i, slice[i])
	}

	// 按字符进行遍历
	for index, value := range str {
		fmt.Printf("%d,%c\n", index, value)
	}

	count := 0
	sum := 0

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
			fmt.Printf("%d,%d\n", count, sum)
		}
	}

	n := 20
	fmt.Println(1)
	fmt.Println(2)
	if n > 10 {
		goto test1
	}
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
test1:
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)

}
