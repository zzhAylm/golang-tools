package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChanFunGoroutine(t *testing.T) {
	fmt.Println("zzhAylm")

	var intChan chan int
	intChan = make(chan int, 3)

	intChan <- 10
	intChan <- 20
	intChan <- 30

	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan

	fmt.Printf("%v,%v,%v\n", num1, num2, num3)

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 3)

	var map1 = make(map[string]string, 10)
	var map2 = make(map[string]string, 10)

	map1["zzh1"] = "ylm"
	map1["zzh2"] = "ylm"
	map1["zzh3"] = "ylm"
	map2["zzh1"] = "ylm"
	map2["zzh2"] = "ylm"
	map2["zzh3"] = "ylm"

	mapChan <- map1
	mapChan <- map2

	fmt.Printf("%v,%v\n", <-mapChan, <-mapChan)

	catChan := make(chan CatTest, 10)

	catTest1 := CatTest{
		18, "zzh",
	}
	catTest2 := CatTest{
		18, "zzh",
	}
	catChan <- catTest1
	catChan <- catTest2
	fmt.Printf("%v,%v\n", <-catChan, <-catChan)

	pointChan := make(chan *CatTest, 10)
	point1 := &CatTest{
		18, "zzh",
	}
	point2 := &CatTest{
		18, "zzh",
	}

	pointChan <- point1
	pointChan <- point2
	pointChan <- &catTest1
	pointChan <- &catTest1
	fmt.Printf("%v,%v,%v,%v\n", <-pointChan, <-pointChan, <-pointChan, <-pointChan)

	interfaceChin := make(chan interface{}, 10)

	interfaceChin <- CatTest{
		Age: 19, Name: "zzh",
	}
	interfaceChin <- 10
	interfaceChin <- "zzh"

	cat := <-interfaceChin
	fmt.Printf("%v,%v,%v\n", cat, <-interfaceChin, <-interfaceChin)

	fmt.Printf("%T\n", cat)

	//  类型断言
	newCat := cat.(CatTest)
	fmt.Printf("%v\n", newCat.Name)
	ChanFunGoroutine()

}

type CatTest struct {
	Age  int
	Name string
}

func TestChanFunGoroutine1(t *testing.T) {

	intChan := make(chan int, 10)

	intChan <- 10
	intChan <- 10
	intChan <- 10
	intChan <- 10
	close(intChan)
	// 管道关闭之后，不可以进行放入了，但是可以取出数据
	//intChan <- 10
	fmt.Println(<-intChan, <-intChan, <-intChan, <-intChan)

	// 管道的遍历
	numChan := make(chan int, 100)

	for i := 0; i < 100; i++ {
		numChan <- i
	}

	close(numChan)
	for num := range numChan {

		fmt.Println(num)
	}

	ChanFunGoroutine()
}

func TestChanFunGoroutine2(t *testing.T) {

	intChan := make(chan int, 10)
	exitChan := make(chan bool, 10)

	go Write(intChan)
	go Read(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if ok {
			fmt.Println("结束")
			break
		}
	}
	ChanFunGoroutine()
}

// 如果一个管道只有写没有读，会直接报错

// 读写的速度可以是不一样的，他会自动平衡读写的速度关系，因为阻塞的关系

func Write(intChin chan int) {
	for i := 0; i < 20; i++ {
		fmt.Println("写入数据=", i)
		intChin <- i
	}
	close(intChin)
}
func Read(intChin chan int, exitChan chan bool) {
	for v := range intChin {
		fmt.Println("读取数据=", v)
	}
	exitChan <- true
	close(exitChan)
}

func TestChanFunGoroutine3(t *testing.T) {

	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	for i := 0; i < 1000; i++ {
		intChan <- i
	}
	close(intChan)
	for i := 0; i < 4; i++ {
		go addSSChan(intChan, resultChan, exitChan)
	}

	for i := 0; i < 4; i++ {
		_, ok := <-exitChan
		if ok {
			fmt.Println("一个子程序完成")
		}
	}

	//遍历管道是，需要先关闭管道，然后在遍历，不然遍历到最后会报错
	close(exitChan)
	close(resultChan)
	for num := range resultChan {
		fmt.Println("奇数", num)
	}
	fmt.Println("主程序退出")
	ChanFunGoroutine()
}
func addSSChan(intChan chan int, resultChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			fmt.Println("没有获取到数据，推出")
			exitChan <- true
			break
		}
		if num%2 != 0 {
			resultChan <- num
		}
	}
}

func TestChanFunGoroutine4(t *testing.T) {
	// 管道可以声明为只读或者只写

	// 默认 ： 可读可写
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//只写
	var chan3 chan<- int
	chan3 = make(chan int, 10)
	for i := 0; i < 10; i++ {
		chan3 <- i
	}
	//只读
	var chan4 <-chan int
	chan4 = make(chan int, 10)
	for num := range chan4 {
		fmt.Println(num)
	}
}

func TestChanFunGoroutine5(t *testing.T) {
	// 管道可以声明为只读或者只写

	// 默认 ： 可读可写
	intChan1 := make(chan int, 10)
	intChan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan1 <- i
	}
	for i := 0; i < 10; i++ {
		intChan2 <- i
	}

	for {
		select {
		case v := <-intChan1:
			fmt.Println("chan1=", v)
		case v := <-intChan2:
			fmt.Println("chan2=", v)
		default:
			fmt.Println("都没取到")
			return
		}
	}

}

// 错误的捕获
func TestChanFunGoroutine6(t *testing.T) {
	// 管道可以声明为只读或者只
	// 子协程报错，不要影响主线程的执行
	go mapTest()
	go sayHello()
	time.Sleep(time.Second)
	fmt.Println("主线程")

}

func mapTest() {
	defer func() { // 捕获错误
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var maps map[int]string
	maps[0] = "zzh" // error
	fmt.Println("子线程")
}
func sayHello() {

	for i := 0; i < 10; i++ {
		fmt.Println("say hello")
	}

}
