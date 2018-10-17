package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

const (
	ServerNetworkType = "tcp4"
	ServerAddr        = "127.0.0.1:8080"
	Delimiter         = '\t'
)

func read(conn net.Conn) ([]byte, error) {
	readBytes := make([]byte, 1) // 防止从conn中读取多余的数据，每读取一个字符都要检查该字符是否是分界符号
	var buf bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return nil, err
		}
		if readBytes[0] == Delimiter { // 如果接收的数据是数据块的分隔符,忽略
			break
		}
		buf.Write(readBytes)
	}
	return buf.Bytes(), nil
}

func write(conn net.Conn, content []byte) (int, error) {
	var buffer bytes.Buffer
	buffer.Write(content)
	buffer.WriteByte(Delimiter)
	return conn.Write(buffer.Bytes())
}

//------------------------------------ server -----------------------------------------
func serverHandleConn(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // 关闭闲置连接的功能，释放资源。根据场景可用可不用
		readBytes, err := read(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("The connect had been closed")
			} else {
				fmt.Println("Read error:", err.Error())
			}
			break
		}
		fmt.Println("    ", strings.Split(conn.RemoteAddr().String(), ":")[0]+":", string(readBytes))

		if len(readBytes) > 4 {
			pong := []byte("pong" + string(readBytes[4:]))
			_, err := write(conn, pong)
			if err != nil {
				break
			}
		}
	}
}

func serverRun() {
	listener, err := net.Listen(ServerNetworkType, ServerAddr)
	if err != nil {
		fmt.Println("Listen error: ", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("服务已经启动，port:8080，等待客户端连接......")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error: ", err.Error())
			continue
		}
		fmt.Println("接受客户端", conn.RemoteAddr(), "连接")
		go serverHandleConn(conn)
	}
}

//------------------------------------client-------------------------------------------

func clientHandleConn(conn net.Conn) {
	defer conn.Close()

	go func() { // 接收处理数据
		for {
			rd, err := read(conn)
			if err != nil {
				break
			}
			fmt.Println("    ", string(rd))
		}
	}()

	cnt := 0
	for { // 循环发送数据
		if _, err := write(conn, []byte(fmt.Sprintf("ping_NO.%d", cnt))); err != nil {
			break
		}

		cnt++

		time.Sleep(time.Second)
		if cnt > 99 {
			break
		}
	}
}

func clientRun() {
	conn, err := net.Dial(ServerNetworkType, ServerAddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Connect server %v success\n", conn.RemoteAddr())
	clientHandleConn(conn)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: <%s> <client|server>\n", os.Args[0])
		return
	}

	switch os.Args[1] {
	case "client":
		clientRun()
	case "server":
		serverRun()
	default:
		fmt.Printf("para is wrong: <%s> <client|server>\n", os.Args[0])
		return
	}
}
