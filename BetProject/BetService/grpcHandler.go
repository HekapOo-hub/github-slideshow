package main

import (
	pb "BetService/Bet"
	"context"
)

type GRPCHandler struct {
	a *AuthenticationService
	pb.UnimplementedBetServiceServer
}

func (h *GRPCHandler) MakeBet(ctx context.Context, in *pb.BetWithToken) (*pb.Response, error) {
	id, err := h.a.CheckRole([]string{"user"}, in.Token)
	if err != nil {
		return nil, err
	}
	if id == -1 {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	str, err := h.a.MakeBet(&Bet{CompetitionId: int(in.Bet.CompetitionId),
		UserId: id, Result: in.Bet.Result,
		Money: int(in.Bet.Money),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
func (h *GRPCHandler) RemoveBet(ctx context.Context, in *pb.IdWithToken) (*pb.Response, error) {
	id, err := h.a.CheckRole([]string{"bookmaker"}, in.Token)
	if err != nil {
		return nil, err
	}
	if id == -1 {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	str, err := h.a.RemoveBet(int(in.CId), int(in.UId))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
func (h *GRPCHandler) GetAll(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	id, err := h.a.CheckRole([]string{"bookmaker", "admin"}, in.Token)
	if err != nil {
		return nil, err
	}
	if id == -1 {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	bets, err := h.a.GetAll()
	if err != nil {
		return nil, err
	}
	all := "All bets:\n"
	for _, bet := range bets {
		all += bet.String()
	}
	return &pb.Response{Response: all}, nil
}
func (h *GRPCHandler) GetMyBets(ctx context.Context, in *pb.IdWithToken) (*pb.Response, error) {
	id, err := h.a.CheckRole([]string{"user"}, in.Token)
	if err != nil {
		return nil, err
	}
	if id == -1 {
		return &pb.Response{Response: "you don't have access to this operation"}, nil
	}
	bets, err := h.a.repository.GetMyBets(id)
	if err != nil {
		return nil, err
	}
	all := "Your bets:\n"
	for _, bet := range bets {
		all += bet.String()
	}
	return &pb.Response{Response: all}, nil
}
