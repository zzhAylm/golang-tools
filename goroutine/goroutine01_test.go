package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	go Test() // 开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main hello golang!", i)
		time.Sleep(time.Second)
	}
}

func TestCPUTest(t *testing.T) {
	CPUTest()

}
