package main

import (
	"fmt"
)

// 数组和切片
func main() {

	int0 := make([]int, 0, 5)

	int1 := append(int0, 1)

	printSlice(int0)
	printSlice(int1)

	//strings.HasSuffix()

}

func printSlice(slice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
