package main

import (
	"flag"
	pb "goNotes/grpc/proto/loginProto"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const serverAddress = ":50051"

func main() {
	var name, pwd string

	flag.StringVar(&name, "name", "", "user name")
	flag.StringVar(&pwd, "pwd", "", "user password")
	flag.Parse()
	if name == "" || pwd == "" {
		log.Printf("name or pwd must has value, usage: <program file> --name=<username> --pwd=<password>")
		return
	}

	// 连接grpc服务器
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 实例化对象loginerClient，对象实现接口LoginerClient
	client := pb.NewLoginerClient(conn)

	// 调用方法，输入参数LoginRequest对象，得到结果LoginReply对象
	r, err := client.Login(context.Background(), &pb.LoginRequest{Username: name, Password: pwd})
	if err != nil {
		log.Fatalf("could not get result: %v", err)
	}
	log.Printf("Message: %+v", r)
}
