package main

import (
	pb "BetService/Bet"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":40042"
)

func main() {
	s := grpc.NewServer()

	h := &GRPCHandler{&AuthenticationService{NewRepository()}, pb.UnimplementedBetServiceServer{}}
	pb.RegisterBetServiceServer(s, h)
	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}

}
