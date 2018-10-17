package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	ServerNetworkType = "udp4"
	ServerAddr        = "127.0.0.1:8080"
)

type OnlineUser struct {
	User       map[string]*net.UDPAddr // 保存用户名
	LatestTime map[string]time.Time    // 保存最新的时间
	mutex      sync.Mutex              // 同步
}

func NewOnlineUser() *OnlineUser {
	return &OnlineUser{User: make(map[string]*net.UDPAddr), LatestTime: make(map[string]time.Time)}
}

func (ou *OnlineUser) Add(userName string, udpAddr *net.UDPAddr) bool {
	ou.mutex.Lock()
	defer ou.mutex.Unlock()
	if _, ok := ou.User[userName]; ok {
		return false
	}
	ou.User[userName] = udpAddr
	ou.LatestTime[userName] = time.Now()
	return true
}

func (ou *OnlineUser) Del(userName string) {
	ou.mutex.Lock()
	delete(ou.User, userName)
	delete(ou.LatestTime, userName)
	ou.mutex.Unlock()
}

func (ou *OnlineUser) GetAllUserName() []string {
	ou.mutex.Lock()
	defer ou.mutex.Unlock()
	userNames := make([]string, len(ou.User))
	cnt := 0
	for un, _ := range ou.User {
		userNames[cnt] = un
		cnt++
	}
	return userNames
}

func (ou *OnlineUser) Updatetime(udpAddr *net.UDPAddr) bool {
	userName, ok := ou.GetUserName(udpAddr)

	ou.mutex.Lock()
	if ok {
		ou.LatestTime[userName] = time.Now()
	}
	ou.mutex.Unlock()
	return ok
}

func (ou *OnlineUser) onlineCheck() []string {
	userNames := ou.GetAllUserName()
	var timeoutName []string
	now := time.Now()
	ou.mutex.Lock()
	for _, userName := range userNames {
		if now.Sub(ou.LatestTime[userName]).Seconds() > 5 {
			timeoutName = append(timeoutName, userName)
		}
		ou.LatestTime[userName] = now
	}
	ou.mutex.Unlock()

	for i := 0; i < len(timeoutName); i++ {
		ou.Del(timeoutName[i])
	}
	return timeoutName
}

func (ou *OnlineUser) GetUserName(udpAddr *net.UDPAddr) (string, bool) {
	ou.mutex.Lock()
	defer ou.mutex.Unlock()
	ok := false
	var uN string
	for username, ua := range ou.User {
		if ua.String() == udpAddr.String() {
			uN = username
			ok = true
			break
		}
	}
	return uN, ok
}

func (ou *OnlineUser) SendAllUser(udpConn *net.UDPConn, msg []byte) {
	cnt := 0
	udpAddrs := make([]*net.UDPAddr, len(ou.User))
	ou.mutex.Lock()
	for _, udpAddr := range ou.User {
		udpAddrs[cnt] = udpAddr
		cnt++
	}
	ou.mutex.Unlock()

	for _, udpAddr := range udpAddrs {
		go func(udpAddr *net.UDPAddr) {
			if udpAddr != nil {
				udpConn.WriteToUDP(msg, udpAddr)
			}
		}(udpAddr)
	}
}

func checkOnlineHandle(udpConn *net.UDPConn, ou *OnlineUser) {
	ticker := time.NewTicker(5 * time.Second)
	var toNames []string
	for {
		<-ticker.C
		toNames = ou.onlineCheck()
		if len(toNames) > 0 {
			msg := "\n---------- " + fmt.Sprintf("%v", toNames) + " get away. ----------\n"
			fmt.Println(msg)
			ou.SendAllUser(udpConn, []byte(msg))
		}
		//		fmt.Println("test-------", ou.GetAllUserName())
	}
}

//------------------------------------ server --------------------------------------------

func handleMsg(udpConn *net.UDPConn, msg []byte, udpAddr *net.UDPAddr, ou *OnlineUser) {
	cmd := strings.SplitAfterN(string(msg), ":", 2)
	if len(cmd) > 1 {
		switch cmd[0] {
		case "connect:":
			ok := ou.Add(cmd[1], udpAddr)
			if ok {
				msg := "\n++++++++++ " + cmd[1] + " is on line. ++++++++++\n"
				fmt.Println(msg)
				ou.SendAllUser(udpConn, []byte(msg))
			}
		case "heartbeat:":
			ou.Updatetime(udpAddr)
		case "chat:":
			userName, ok := ou.GetUserName(udpAddr)
			if ok {
				msgData := userName + ": " + cmd[1]
				fmt.Println(msgData)
				ou.SendAllUser(udpConn, []byte(msgData))
			}
		case "users:":
			aun := ou.GetAllUserName()
			var usernameMsg string
			for k, name := range aun {
				usernameMsg += fmt.Sprintf("    %d: %s", k+1, name+"\n")
			}
			udpConn.WriteToUDP([]byte(usernameMsg), udpAddr)
		default:
			fmt.Println("无效命令")
		}
	}
}

func severHandleConn(udpConn *net.UDPConn, ou *OnlineUser) {
	bufRD := make([]byte, 1024)
	n, udpAddr, err := udpConn.ReadFromUDP(bufRD)
	if err != nil {
		fmt.Println("readFromUDP error:", err.Error())
		os.Exit(1)
	}

	if n > 0 {
		go handleMsg(udpConn, bufRD[:n], udpAddr, ou)
	}
}

func serverRun() {
	// 监听地址
	udpAddr, err := net.ResolveUDPAddr(ServerNetworkType, ServerAddr)
	if err != nil {
		return
	}
	// 监听连接
	udpConn, err := net.ListenUDP(ServerNetworkType, udpAddr)
	if err != nil {
		fmt.Println("Listen error:", err.Error())
		return
	}
	defer udpConn.Close()

	fmt.Println("udp server running......")

	ou := NewOnlineUser()
	go checkOnlineHandle(udpConn, ou)
	for {
		severHandleConn(udpConn, ou)
	}
}

//------------------------------------client-------------------------------------------

func sendHeartbeat(udpConn net.Conn) {
	ticker := time.NewTicker(3 * time.Second)
	for {
		<-ticker.C
		_, err := udpConn.Write([]byte("heartbeat:" + udpConn.LocalAddr().String()))
		if err != nil {
			fmt.Println("write error:", err.Error())
			os.Exit(1)
		}
	}
}

func clientHandleConn(udpConn net.Conn) {
	bufRD := make([]byte, 1024)
	go func() {
		for {
			n, err := udpConn.Read(bufRD)
			if err != nil {
				fmt.Println("read error:", err.Error())
				os.Exit(-1)
			}
			fmt.Println(string(bufRD[:n]), "\n")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	var cmd string
	for {
		if scanner.Scan() {
			if scanner.Text() == "users" {
				cmd = "users:"
			} else {
				cmd = "chat:"
			}
			_, err := udpConn.Write([]byte(cmd + scanner.Text()))
			if err != nil {
				fmt.Println("send error:", err.Error())
				return
			}
			fmt.Println()
		}
	}
}

func clientRun() {
	// 监听地址
	udpAddr, err := net.ResolveUDPAddr(ServerNetworkType, ServerAddr)
	if err != nil {
		return
	}
	//	// 监听连接
	//	udpConn, err := net.ListenUDP(ServerNetworkType, udpAddr)
	//	if err != nil {
	//		fmt.Println("Listen error:", err.Error())
	//		return
	//	}
	udpConn, err := net.DialUDP(ServerNetworkType, nil, udpAddr)
	if err != nil {
		fmt.Println("Dial error:", err.Error())
		return
	}
	defer udpConn.Close()

	fmt.Printf("请输入用户名称：")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		_, err := udpConn.Write([]byte("connect:" + scanner.Text()))
		if err != nil {
			fmt.Println("write error:", err.Error())
		}
	}
	go sendHeartbeat(udpConn)

	clientHandleConn(udpConn)
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
