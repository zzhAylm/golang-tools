package main

import (
	"encoding/json"
	"fmt"
)

// 封装，继承和多态
// 封装：对属性的封装，默认类型，无法访问，只能通过方法访问器属性值
// 继承：代码复用抽象出公共属性
func main() {
	var son1 Son1
	son1.Age = 18
	son1.Name = "zzh"
	son1.Major = "zh"

	marshal, err := json.Marshal(son1)
	if err == nil {
		fmt.Println(string(marshal))
	}

}

type Son1 struct {
	Father        // 继承father类的所有属性
	Major  string `json:"major"` //专业
}
type Son2 struct {
	Father
	Grade string `json:"grade"` //专业
}
type Father struct {
	Name string `json:"name"` //专业
	Age  int    `json:"age"`  //专业
}
