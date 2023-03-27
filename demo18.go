package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"main/morestrings"
)

//go build  ; 编译
//go mod tidy ：导入需要的包，并且删除不再使用的模块

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
