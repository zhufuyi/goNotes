package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

const (
	ServerNetworkType = "tcp4"
	ServerAddr        = "192.168.8.104:8080"
	Delimiter         = '\t' // 客户端和服务器端约定的数据块边界为‘\t’
)

func read(bufReader *bufio.Reader) ([]byte, error) {
	readBytes, err := bufReader.ReadBytes(Delimiter) // 直到读取到分界符号返回数据，包括分割符号
	if err != nil {
		return nil, err
	}
	return readBytes[:len(readBytes)-1], nil
}

func write(writeBuf *bufio.Writer, content []byte) (int, error) {
	writeBuf.Write(content)
	writeBuf.WriteByte(Delimiter)
	size := writeBuf.Buffered() // 获得缓存的字节数
	err := writeBuf.Flush()
	return size, err
}

//------------------------------------ server --------------------------------------------

func serverHandleConn(conn net.Conn) {
	defer conn.Close()

	bufReader := bufio.NewReader(conn) // 新建的缓冲读取器不能放在for循环的里边，否则会造成数据块被遗漏
	writeBuf := bufio.NewWriter(conn)
	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // 关闭闲置连接的功能，释放资源。根据场景可用可不用
		readBytes, err := read(bufReader)
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
			_, err := write(writeBuf, pong)
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
	bufReader := bufio.NewReader(conn) // 新建的缓冲读取器不能放在for循环的里边，否则会造成数据块被遗漏
	writeBuf := bufio.NewWriter(conn)

	go func() { // 接收处理数据
		for {
			rd, err := read(bufReader)
			if err != nil {
				break
			}
			fmt.Println("    ", string(rd))
		}
	}()

	cnt := 0
	for { // 循环发送数据
		if _, err := write(writeBuf, []byte(fmt.Sprintf("ping_NO.%d", cnt))); err != nil {
			break
		}

		cnt++
		if cnt%5 == 0 {
			time.Sleep(5 * time.Second)
		}
		if cnt == 100 {
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
