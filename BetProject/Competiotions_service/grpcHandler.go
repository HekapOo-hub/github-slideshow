package main

import (
	pb "Competiotions_service/competition"
	"context"
)

type GRPCHandler struct {
	a *AuthenticationService
	pb.UnimplementedCompetitionServiceServer
}

func (h *GRPCHandler) CreateCompetition(ctx context.Context, in *pb.Competition) (*pb.Response, error) {
	ok, err := h.a.CheckRole([]string{"admin"}, in.Token)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	str, err := h.a.Create(&Competition{Name: in.Name, Result: in.Result})
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
func (h *GRPCHandler) GetById(ctx context.Context, in *pb.Id) (*pb.Response, error) {
	ok, err := h.a.CheckRole([]string{"user", "admin", "bookmaker"}, in.Token)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	c, err := h.a.GetById(int(in.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: c.String()}, nil
}
func (h *GRPCHandler) GetAll(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	ok, err := h.a.CheckRole([]string{"user", "admin", "bookmaker"}, in.Token)
	if err != nil {
		return nil, err
	}

	if !ok {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	comps, err := h.a.GetAll()
	if err != nil {
		return nil, err
	}
	all := "All competitions:\n"
	for _, comp := range comps {
		all += comp.String()
	}
	return &pb.Response{Response: all}, nil
}
func (h *GRPCHandler) SetResult(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	ok, err := h.a.CheckRole([]string{"admin"}, in.Token)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	str, err := h.a.SetResult(int(in.Id), in.Result)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
