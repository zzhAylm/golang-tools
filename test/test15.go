package main

import (
	"encoding/json"
	"fmt"
)

// 结构体

type Cat struct {
	Age   int
	Name  string
	Color string
	arr   [10]int
	ptr   *int              //默认是nil
	slice []string          //默认是nil ，赋值之前需要先进行make
	maps  map[string]string //默认是nil 赋值之前需要先进行make
}

// 基本数据类型、数组、引用类型
// 值类型： 基本数据类型（int，float，bool，string）数组和结构体
// 引用类型： 指针，slice切片，map，管道chan，interface都是引用类型
func main() {

	var cat Cat
	cat.Name = "zzh"
	cat.Age = 18
	cat.Color = "白色"
	fmt.Println(cat)

	ints := 100
	cat.ptr = &ints

	fmt.Println(*cat.ptr)
	cat.slice = make([]string, 10)
	cat.slice[0] = "zzh"

	cat.maps = make(map[string]string, 10)
	cat.maps["zzh"] = "zzh"
	fmt.Println(cat)

	// 创建结构体实例的方式
	cat2 := Cat{
		18, "zzh", "红色", [10]int{10}, nil, make([]string, 10), make(map[string]string, 10),
	}

	fmt.Println(cat2)

	var cat3 = new(Cat)

	// (*指针).属性=值
	(*cat3).Name = "zzh"
	// 也可以使用，指针.属性=值 ,原因是go，为了开发方便，底层做了优化
	cat3.Age = 18
	fmt.Println(*cat3)

	var cat4 = &Cat{}
	cat4.Age = 18
	(*cat4).Age = 20
	//获取属性也可以直接是  指针.属性 或者是 (*指针).属性
	fmt.Println(cat4.Age)
	fmt.Println(*cat4)

	//值类型： 基本数据类型（int，float，bool，string）数组和结构体
	//引用类型： 指针，slice切片，map，管道chan，interface都是引用类型

	// 结构体中的属性，地址是连续存在的，如果是指针类型，指针类型本身的地址是连续的，他们 指向的地址不一定是连续的

	//	结构体的拷贝：需要两个结构体的数据行和类型完全一样，（字段的名字、类型、个数、完全一样），
	//	取的别名和原来的数据类型属于两种数据类型，他认为是一种新的数据类型

	stu1 := Student{200, "zzh"}

	fmt.Println(stu1)

	marshal, err := json.Marshal(stu1)

	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println(string(marshal))
}

// Student tag ：结构体序列化时，json中的属性名字自定义
type Student struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
