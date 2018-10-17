package main

import (
	"log"
	"net"

	pb "goNotes/grpc/proto/loginProto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	host = ":50051"
)

type server struct {
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	rpl := &pb.LoginReply{}

	if in.Username == "grpc" && in.Password == "123456" {
		rpl.Status = true
		rpl.ErrorMsg = ""
	} else {
		rpl.Status = false
		rpl.ErrorMsg = "failed to verify"
	}

	return rpl, nil
}

func main() {
	// 监听服务
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("rpc server listen on %s", host)

	// 实例化对象
	s := grpc.NewServer()
	// 注册方法
	pb.RegisterLoginerServer(s, &server{})
	// 启动gppc服务器
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
