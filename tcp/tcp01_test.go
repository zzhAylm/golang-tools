package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"testing"
)

// server
func TestFunTcp01(t *testing.T) {

	// tcp 协议：长连接
	// http 协议：短链接

	// telnet 断开连接
	// ctrl + ]  输入quit

	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("listen,", listen)
	defer listen.Close()
	for {
		accept, errAccept := listen.Accept()
		if errAccept != nil {
			fmt.Println("连接失败", err)
		} else {
			fmt.Println("连接成功", accept.LocalAddr().Network(), accept.LocalAddr().String())
		}
		// 起一个协程，处理客户端请求
		go process(accept)
	}

	FunTcp()
}
func process(con net.Conn) {
	defer con.Close()
	for {
		bytes := make([]byte, 1024)
		// 等待客户端通过conn发送消息
		// 等待客户端发送数组，阻塞式的
		fmt.Printf("等待客户端输入数据 %s\n", con.RemoteAddr().String())
		n, err := con.Read(bytes)
		if err != nil {
			fmt.Println("客户端退出,断开连接", err)
			return
		}
		// 打印客户端发送的信息
		fmt.Print("服务端接收的消息：" + string(bytes[:n]))
	}

}

// client
func TestFunTcp02(t *testing.T) {

	dial, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	fmt.Println("连接成功，请输入要发送的数据")
	defer dial.Close()
	for {

		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取控制台输入的数据失败")
			return
		}

		if readString == "exit\n" {
			fmt.Println("停止输入，断开连接")
			return
		}
		n, err := dial.Write([]byte(readString))
		if err != nil {
			fmt.Println("发送给服务端失败，断开连接")
			return
		}
		fmt.Printf("客户端发送成功，message is %s nubmer is %d", readString, n)
	}

}
