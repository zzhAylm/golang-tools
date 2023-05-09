package reflex

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFunReflex1(t *testing.T) {
	// 反射的基本介绍：

	FunReflex()

}
func TestFunReflex2(t *testing.T) {
	// 反射的基本介绍：
	num1 := 10
	ReflexFun(num1)
	reflex := Reflex{Name: "zzh", Age: 18}
	ReflexFun(reflex)
	FunReflex()
}

type Reflex struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func ReflexFun(data interface{}) {
	t := reflect.TypeOf(data)
	fmt.Println(t)

	of := reflect.ValueOf(data)
	switch value := of.Interface().(type) {
	case int:
		fmt.Println("int", value)
	case Reflex:
		fmt.Println("reflex", of.NumField(), value)
	default:
		fmt.Println("other", value)
	}

}

func TestFunReflex3(t *testing.T) {
	// 反射的基本介绍：

	FunReflex()

}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {

}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {

	return true
}

func TestFunReflex4(t *testing.T) {
	// 反射的基本介绍：
	monster := Monster{
		Age:     18,
		Name:    "zzh",
		Source:  100,
		Address: "bj",
	}
	valueOf := reflect.ValueOf(monster)
	typeOf := reflect.TypeOf(monster)
	field := typeOf.NumField()
	numField := valueOf.NumField()
	fmt.Println(field)
	fmt.Println(numField)
	for i := 0; i < numField; i++ {
		value := valueOf.Field(i)
		fmt.Println(i, value)
		structField := typeOf.Field(i)
		fmt.Println(structField.Name, structField.Tag.Get("json"))
	}

	methodNum := typeOf.NumMethod()
	methodNum1 := valueOf.NumMethod()

	fmt.Println("method num is ", methodNum)
	fmt.Println("method num is ", methodNum1)

	for i := 0; i < methodNum; i++ {
		method := typeOf.Method(i)
		fmt.Println(method.Name)

	}
	valueMethod := valueOf.Method(0)
	valueMethod.Call(nil)

	var params []reflect.Value
	param1 := append(params, reflect.ValueOf(20))
	valueOf.Method(1).Call(param1)

	param2 := append(param1, reflect.ValueOf(10))

	call := valueOf.Method(2).Call(param2)

	fmt.Println(call)
	FunReflex()

}

type Monster struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Source  int    `json:"source"`
	Address string `json:"address"`
}

func (m Monster) Method1() {
	fmt.Println("do method1")

}
func (m Monster) Method2(n int) {
	fmt.Println(n)

}
func (m Monster) Method3(x, y int) {
	fmt.Println(x + y)
}
