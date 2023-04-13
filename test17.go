package main

import "fmt"

type Zzh struct {
	age  int
	name string
}

// 方法，且只有 Zzh 类型可以访问 ,(变量 结构体类型)相当于穿参，传一个Zzh结构体类型参数，只能穿入Zzh结构类型
func (z Zzh) fun1(i int) int {
	fmt.Println(z.age, z.name)
	return 100 + i
}

// 不是指针传递
func (z *Zzh) fun2(i int) int {
	fmt.Println(z.age, z.name)
	return 100 + i
}

// 方法， 结构体是值传递 ，方法和结构体绑定
func main() {

	var z Zzh
	z.name = "zzh"
	z.age = 18

	// 为某个type绑定一个方法
	fmt.Println(z.fun1(7))

	var z2 Zzh
	z2.age = 18
	z2.name = "ylm" // z2的一个副本传递到 fun1 方法中，不会对z2造成影响 ，和函数传递中的结构体参数是值传递一样
	fmt.Println(z2.fun1(7))

	var z3 Zzh
	z3.age = 18
	z3.name = "zzhAylm"
	// 正常写法 （&z3）.fun2()  ,编译对我们的代码进行了优化可以 直接  对象.fun  ，编译器会自定帮我加入 （&对象).fun ,帮我们获取指针
	// 仅以多使用指针类型，这样可以减少值传递过程中的拷贝过程
	z3.fun2(7)

	var i Integer = 10
	i.funInt()
	fmt.Println(i)

	var s = Stu{"zzh", 18}
	fmt.Println(s)
	fmt.Println(&s)

	var stu1 = Stu{"小明", 18}
	stu2 := Stu{"小明", 18}
	stu3 := Stu{Name: "小明", Age: 18}

	stu4 := &Stu{Name: "小明", Age: 18}
	var stu5 = &Stu{Name: "小明", Age: 18}

	fmt.Println(stu1, stu2, stu3, stu4, stu5)
}

type Integer int

func (i *Integer) funInt() {
	*i = 100
	fmt.Println(*i)
}

type Stu struct {
	Name string
	Age  int
}

func (stu *Stu) String() string {
	return fmt.Sprintf("name=[%s],age=[%d]", stu.Name, stu.Age)
}
