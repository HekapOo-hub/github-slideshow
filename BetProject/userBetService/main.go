package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "userBetService/user"
)

const (
	port = ":40043"
)

func main() {
	s := grpc.NewServer()

	h := &GRPCHandler{&AuthWithTokenService{&AuthorizationService{NewRepository()}}, pb.UnimplementedUserServiceServer{}}
	
	pb.RegisterUserServiceServer(s, h)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
