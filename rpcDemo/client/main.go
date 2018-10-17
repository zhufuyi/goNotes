package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"goNotes/rpcDemo"
)

var rpcClient *rpc.Client

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	rpcClient = jsonrpc.NewClient(conn)

	args := rpcDemo.Args{3, 4}
	fmt.Println(callAdd(args))
	fmt.Println(callSub(args))
	fmt.Println(callMul(args))
	fmt.Println(callDiv(args))
}

func callAdd(args rpcDemo.Args) (int, error) {
	var result int
	err := rpcClient.Call(rpcDemo.Add, args, &result)
	return result, err
}

func callSub(args rpcDemo.Args) (int, error) {
	var result int
	err := rpcClient.Call(rpcDemo.Sub, args, &result)
	return result, err
}

func callMul(args rpcDemo.Args) (int, error) {
	var result int
	err := rpcClient.Call(rpcDemo.Mul, args, &result)
	return result, err
}

func callDiv(args rpcDemo.Args) (float64, error) {
	var result float64
	err := rpcClient.Call(rpcDemo.Div, args, &result)
	return result, err
}
