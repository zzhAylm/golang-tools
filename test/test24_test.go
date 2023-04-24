package main_test

import "testing"

// testing 框架

// TestAdd testing 框架的使用
//
//	go 文件要以 xxx_test.go  为结尾， test函数要以  Test（要测试的函数名且第一个字母要大写）
func TestAdd(t *testing.T) {
	res := Add(10, 20)
	t.Logf("执行成功 Add result%v", res)
}
func Add(a, b int) int {
	return a + b
}
