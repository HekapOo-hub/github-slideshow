package main

import (
	pb2 "BetClientAPI/Bet"
	pb1 "BetClientAPI/competition"
	pb3 "BetClientAPI/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

const (
	port     = ":40040"
	address1 = "localhost:40041"
	address2 = "localhost:40042"
	address3 = "localhost:40043"
)

func main() {
	conn1, err := grpc.Dial(address1, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Не могу подключиться: %v", err)
	}
	defer conn1.Close()

	client1 := pb1.NewCompetitionServiceClient(conn1)

	conn2, err := grpc.Dial(address2, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Не могу подключиться: %v", err)
	}
	defer conn2.Close()

	client2 := pb2.NewBetServiceClient(conn2)
	conn3, err := grpc.Dial(address3, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Не могу подключиться: %v", err)
	}
	defer conn2.Close()
	client3 := pb3.NewUserServiceClient(conn3)
	server := &EchoServer{echo.New(), client1, client2, client3}
	server.Register()
	err = server.Echo.Start(port)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
}
