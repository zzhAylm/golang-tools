package main

import "fmt"

// 流程控制
func main() {

	var var1 = 10
	var var2 = 20

	if var1 == 10 && var2 > 10 {
		fmt.Printf("var1=%d,var2=%d\n", var1, var2)
	}

	switch var1 {
	case 10:
		fmt.Printf("var=%d\n", var1)
		//fallthrough // 穿透，不管下一个条件满不满足都会执行下一个
	default:
		fmt.Printf("var=%d\n", var1)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
	}

}
