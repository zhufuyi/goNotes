package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"goNotes/rpcDemo"
)

func main() {
	rpc.Register(&rpcDemo.DemoService{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Printf("rpcDemo server listen on %s\n", ":8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			println("accept errors", err.Error())
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
