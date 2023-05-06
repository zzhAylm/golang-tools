package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func Test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test hello world!", i)
		time.Sleep(time.Second)
	}
}

func CPUTest() {
	fmt.Println("print function.")
	fmt.Println("cpu number is ", runtime.NumCPU())
	//

	fmt.Println("set cpu run max number", runtime.GOMAXPROCS(1))
}
