package main

import "fmt"

// 向上转型可以，但是向下转型不可以直接转，需要用到类型断言   接口.(实现结构体)=结构体的类型变量
func main() {

	var myInt interface{}
	var b1 float32 = 10.1
	myInt = b1

	//if err {
	//	fmt.Println("success", b2)
	//}
	//fmt.Println("fail")

	if b2, ok := myInt.(float64); ok {
		fmt.Println("success", b2)
	}
	fmt.Println("fail")

	n1 := 10
	n2 := "str"
	n3 := false
	n4 := S{
		Age:  18,
		Name: "zzh",
	}
	n5 := &S{
		Age:  18,
		Name: "zzh",
	}
	JudgeType(n1, n2, n3, n4, n5)

}

type S struct {
	Age  int
	Name string
}

// JudgeType 判断变量的类型
func JudgeType(items ...interface{}) {

	for _, item := range items {
		switch item.(type) {
		case string:
			fmt.Println("string")
		case float64:
			fmt.Println("float64")
		case float32:
			fmt.Println("float32")
		case int, int32, int64:
			fmt.Println("int,int64")
		case bool:
			fmt.Println("bool")
		case byte:
			fmt.Println("byte")
		case S:
			fmt.Println("S")
		case *S:
			fmt.Println("*S")
		default:
			fmt.Println("其他类型")
		}
	}
}
