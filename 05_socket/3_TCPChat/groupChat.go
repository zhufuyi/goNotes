package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const (
	ServerNetworkType = "tcp4"
	ServerAddr        = "127.0.0.1:8080"
)

type OnlineUser struct {
	User  map[string]net.Conn
	mutex sync.Mutex
}

func NewOnlineUser() *OnlineUser {
	return &OnlineUser{User: make(map[string]net.Conn)}
}

func (ou *OnlineUser) Add(conn net.Conn, bufRD []byte) (string, bool) {
	var userName string
	n, _ := conn.Read(bufRD) // 第一次读取用户名字
	if n == 0 {
		userName = conn.RemoteAddr().Network()
	} else {
		userName = string(bufRD[:n])
	}
	ou.mutex.Lock()
	defer ou.mutex.Unlock()
	if _, ok := ou.User[userName]; ok {
		return userName, false
	}
	ou.User[userName] = conn
	return userName, true
}

func (ou *OnlineUser) Del(userName string) {
	ou.mutex.Lock()
	delete(ou.User, userName)
	ou.mutex.Unlock()
}

func (ou *OnlineUser) SendAllUser(userName string, msg []byte) {
	sum := 0
	conns := make([]net.Conn, len(ou.User))
	ou.mutex.Lock()
	for un, connect := range ou.User {
		if un != userName {
			conns[sum] = connect
			sum++
		}
	}
	ou.mutex.Unlock()

	for _, connect := range conns {
		go func(conn net.Conn) {
			if conn != nil {
				conn.SetWriteDeadline(time.Now().Add(time.Second))
				_, err := conn.Write(msg)
				if err != nil {
					fmt.Println("send error:", err.Error())
				}
			}
		}(connect)
	}
}

//------------------------------------ server --------------------------------------------

func severHandleConn(conn net.Conn, ou *OnlineUser) {
	defer conn.Close()
	bufRD := make([]byte, 1024)
	userName, ok := ou.Add(conn, bufRD)
	if !ok {
		conn.Write([]byte("error: " + userName + " was exist."))
		return
	}
	fmt.Println("\n++++++++++ " + userName + " is online. ++++++++++\n")
	ou.SendAllUser(userName, []byte(userName+" is online"))

	var outMsg string
	for {
		n, err := conn.Read(bufRD)
		if err != nil {
			fmt.Println("\n---------- " + userName + " got away. ----------\n")
			ou.Del(userName)
			ou.SendAllUser(userName, []byte(userName+" got away."))
			break
		}
		outMsg = "[" + userName + "]: " + string(bufRD[:n])
		fmt.Println(outMsg)
		// 转发信息给其他组员
		ou.SendAllUser(userName, []byte(outMsg))
	}
}

func serverRun() {
	netListen, err := net.Listen(ServerNetworkType, ServerAddr)
	if err != nil {
		fmt.Println("Listen error:", err.Error())
		return
	}
	defer netListen.Close()
	ou := NewOnlineUser()
	fmt.Println("Waiting for clients connecting......")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go severHandleConn(conn, ou)
	}
}

//------------------------------------client-------------------------------------------

func clientHandleConn(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("请输入用户名称：")
	inputReader := bufio.NewReader(os.Stdin)
	line, _ := inputReader.ReadBytes('\n')

	_, err := conn.Write(line[:len(line)-1])
	if err != nil {
		fmt.Println("send error:", err.Error())
		return
	}
	bufRD := make([]byte, 1024)
	go func() {
		for {
			n, err := conn.Read(bufRD)
			if err != nil {
				fmt.Println("read error:", err.Error())
				os.Exit(-1)
			}
			fmt.Println(string(bufRD[:n]), "\n")
		}
	}()

	for {
		line, _ := inputReader.ReadBytes('\n')
		_, err := conn.Write(line[:len(line)-1])
		if err != nil {
			fmt.Println("send error:", err.Error())
			return
		}
		fmt.Println()
	}
}

func clientRun() {
	conn, err := net.Dial(ServerNetworkType, ServerAddr)
	if err != nil {
		fmt.Println("connect error:", err.Error())
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
