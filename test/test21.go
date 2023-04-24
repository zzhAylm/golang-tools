package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文件操作
func main() {
	path := "/Users/zzh/Desktop/demo.txt"
	open, err := os.Open(path)
	if err != nil {
		return
	}
	fmt.Println("打开文件成功")

	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			fmt.Println("关闭文件失败", err)
		}
		fmt.Println("关闭文件成功")
	}(open)

	reader := bufio.NewReader(open)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s\n", line)
	}

	// 因为我们没有显示的打开文件，所以也不需要显示的关闭文件，打开和关系的操作放在了ReadFile函数里面
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("读取文件失败")
	}
	fmt.Printf("%s", file)

	//
	//openFile, err := os.OpenFile(path, 1, os.FileMode())
}
