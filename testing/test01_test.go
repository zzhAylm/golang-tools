package testing

import "testing"

// test文件，要以 test01.go （要测试的文件为开头） + _test.go
// 测试函数 Test+要测试的函数名+序号
func TestAdd(t *testing.T) {
	t.Logf("Add函数测试结果=%v", Add(10, 20))
}

func TestUpper(t *testing.T) {
	t.Logf("Upper函数测试结果=%v", Add(10, 20))
}
