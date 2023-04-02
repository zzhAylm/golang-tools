package main

import (
	"fmt"
	"strings"
)

func main() {

	varArr := make([]int, 0, 5)

	// 向切片追加元素
	varArr = append(varArr, 1, 2, 3, 4, 5, 6)

	//len(varArr)： 数组长度 ，cap(varArr)：切片容量
	fmt.Println(varArr, len(varArr), cap(varArr))

	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Println(i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}

	// 对于golang中的字符，ascll中用byte保存，不再acsll中用 int或者更大的数保存
	var c byte = 'c'
	var a byte = 'a'
	var name int = '张'

	fmt.Printf("%c,%c,%c", c, a, name)
}
