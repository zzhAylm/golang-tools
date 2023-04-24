package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// 编译： go build xxx.go   直接执行：xxx.exe
// 直接运行：go run xxx.go
// go注意事项
// go常用转译字符：\t 、\n、\\ 、 \" 、 \r
func main() {
	switch os := runtime.GOOS; os {
	case "mac":
		fmt.Println("OS mac")
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// 没有条件的 switch 和 if else一样
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//
	defer fmt.Println("world")

	//  一个变量可以是 如果赋值了（指向了一个变量）输出就是一个变量
	// 如果一个变量指向一个地址值，则输出就是一个地址值  *(地址值) = 取储地址的值   &(指向一个值的变量）= 获取此变量的指针

	// *p ： 直接修改此地址的值，引用地址传递

	// & 操作符会生成一个指向其操作数的指针
	// Go 拥有指针。指针保存了值的内存地址。
	// 类型 *T 是指向 T 类型值的指针。其零值为 nil。
	// * 操作符表示指针指向的底层值。
	i, j := 42, 2701
	fmt.Println(i)
	fmt.Println(j)
	p := &i // 指向 i
	fmt.Println(p)
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值、

	V := Vertex{1, 2}
	V.Y = 4
	fmt.Println(V)

	//
	v := Vertex{1, 2}
	vv := &v
	fmt.Println(&v)
	fmt.Println(vv)
	fmt.Println(*vv)
	vv.X = 1e9
	fmt.Println(v)

	//
	v1 := Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 := Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 := Vertex{}      // X:0 Y:0
	p1 := &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	var p2 *Vertex
	p2 = &v3
	fmt.Println(v1, v2, v3, p1, p2)
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Vertex 结构体
type Vertex struct {
	X int
	Y int
}
