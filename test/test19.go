package main

import "fmt"

type MyInterface interface {
	Method1()
	Method2() int
	Method3(string) (bool, error)
}

type MyStruct struct{}

func (s MyStruct) Method1() {
	fmt.Println("Method1 called")
}

func (s MyStruct) Method2() int {
	fmt.Println("Method2 called")
	return 42
}

func (s MyStruct) Method3(str string) (bool, error) {
	fmt.Println("Method3 called with string:", str)
	return true, nil
}

// 结构体的方法可以是 结构体类型也可以是结构体指针类型，我们调用的时候直接用对象调用就可以，go语言的编译器会为我们优化
// 过程中如果有两个匿名结构体（父结构体）：且两个结构体有相同的方法或者属性，如果我们使用此方法或者属性，
// 就必须指定匿名结构体，否则编译器不知道使用那个匿名结构体内的方法或者属性就会报错

// 如果结构体内有另一个结构体属性，则这两个是结构体组合

// 建议不要使用多重继承

// Usb interface 接口类型
type Usb interface {
	Method()
	Method2()
	Method3()
}

// Usb2 如果两个接口的方法相同，则结构体实现了usb 也就实现了usb2 ，接口是基于方法实现的
type Usb2 interface {
	Method()
	Method2()
	Method3()
}

type Phone struct {
}

func (p *Phone) Method() {
	fmt.Println(p)
}
func (p *Phone) Method2() {
	fmt.Println(p)
}
func (p *Phone) Method3() {
	fmt.Println(p)
}

func main() {
	var iface MyInterface
	var s MyStruct

	// 多态：
	iface = s

	iface.Method1()
	iface.Method2()
	iface.Method3("hello")

	s.Method1()

	var usb Usb = &Phone{}
	usb.Method()
}
