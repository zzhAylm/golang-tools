package goroutine

import (
	"fmt"
	"testing"
)

// channel 的使用

func TestChanFunChannel(t *testing.T) {

	var intChan chan int
	intChan = make(chan int, 3)

	intChan <- 10

	num := 211
	intChan <- num

	fmt.Printf("intChan %p , 指针本身的地址 %p\n", intChan, &intChan)

	fmt.Printf("channel len=%v , cap =%v\n", len(intChan), cap(intChan))

	num2 := <-intChan
	fmt.Printf("channel num is %d\n", num2)
}
