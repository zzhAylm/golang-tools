package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件copy：
func main() {
	path := "/Users/zzh/Desktop/demo.txt"
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	path1 := "/Users/zzh/Desktop/demo1.txt"
	file1, err1 := os.OpenFile(path1, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		return
	}
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {
			fmt.Println("关闭文件失败")
		}
	}(file1)
	writer := bufio.NewWriter(file1)

	written, err2 := io.Copy(writer, reader)
	if err2 != nil {
		fmt.Println(nil)
	}
	fmt.Println(written)

}
func Add(a, b int) int {
	return a + b
}
