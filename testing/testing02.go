package testing

import (
	"fmt"
	"github.com/goccy/go-json"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (monster *Monster) Save() {
	marshal, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Println(string(marshal))
	path := "/Users/zzh/Company/projects/golang-tools/file/test02.txt"
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}

	// close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	_, err = file.Write(marshal)
	if err != nil {
		return
	}
}

func (monster *Monster) Download() {

	path := "/Users/zzh/Company/projects/golang-tools/file/test02.txt"
	//open, err2 := os.Open(path)
	//if err2 != nil {
	//	fmt.Printf("打开文件 [%s] 失败 \n", path)
	//	return
	//}
	//
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		return
	//	}
	//}(open)
	//
	//reader := bufio.NewReader(open)
	//
	//var data []byte
	//
	//for {
	//	line, _, err2 := reader.ReadBytes()
	//	if err2 == io.EOF {
	//		return
	//	}
	//}

	data, err2 := ioutil.ReadFile(path)
	if err2 != nil {
		fmt.Printf("打开文件 [%s] 失败 \n", path)
		return
	}
	err := json.Unmarshal(data, monster)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
}
