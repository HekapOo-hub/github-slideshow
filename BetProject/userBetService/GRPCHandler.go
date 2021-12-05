package main

import (
	"context"
	pb "userBetService/user"
)

type GRPCHandler struct {
	*AuthWithTokenService
	pb.UnimplementedUserServiceServer
}

func (h *GRPCHandler) CreateUser(ctx context.Context, u *pb.User) (*pb.Response, error) {
	str, err := h.AuthorizationService.CreateUser(&User{Name: u.Name, Password: u.Password,
		Balance: int(u.Balance)})
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
func (h *GRPCHandler) GetAll(ctx context.Context, empty *pb.Empty) (*pb.Response, error) {
	users, err := h.AuthorizationService.GetAll()
	if err != nil {
		return nil, err
	}
	info := "All users:\n"
	for _, user := range users {
		info += user.String()
	}
	return &pb.Response{Response: info}, nil
}
func (h *GRPCHandler) DeleteUser(ctx context.Context, id *pb.Id) (*pb.Response, error) {
	str, err := h.AuthorizationService.DeleteUser(int(id.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
func (h *GRPCHandler) SignIn(ctx context.Context, in *pb.SignInInfo) (*pb.ResponseWithToken, error) {
	token, err := h.AuthWithTokenService.AuthorizeWithToken(in.Name, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.ResponseWithToken{Response: "You entered successfully\n", Token: token}, nil
}
