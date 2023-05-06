package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var resMap = make(map[int]int, 10)
var lock = sync.Mutex{} // 互斥锁

// 线程安全，lock
func TestChannelFun(t *testing.T) {
	for i := 0; i < 20; i++ {
		go jc(i)
	}
	time.Sleep(time.Second)
	fmt.Println(resMap)

}
func jc(n int) int {
	if n <= 1 {
		resMap[n] = n
		return n
	}
	res := 1
	for i := 1; i <= n; i++ {
		res *= n
	}
	lock.Lock()
	resMap[n] = res
	lock.Unlock()
	return res
}
