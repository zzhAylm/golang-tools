package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	path := "/Users/zzh/Desktop/demo.txt"

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("hello world!\n")
	}
	writer.Flush()
}
