package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EchoServer struct {
	*echo.Echo
	repository
}

func (e *EchoServer) makeBet(c echo.Context) error {
	id := c.QueryParam("competitionId")
	cId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result := c.QueryParam("result")
	m := c.QueryParam("money")
	money, err := strconv.Atoi(m)
	if err != nil {
		return err
	}
	str, err := e.repository.MakeBet(&Bet{CompetitionId: cId, Result: result, Money: money})
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, str)
}
func (e *EchoServer) removeBet(c echo.Context) error {
	i := c.QueryParam("cId")
	id, err := strconv.Atoi(i)
	if err != nil {
		return err
	}
	i1 := c.QueryParam("uId")
	id1, err := strconv.Atoi(i1)
	if err != nil {
		return err
	}

	str, err := e.repository.RemoveBet(id, id1)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, str)
}
func (e *EchoServer) getAll(c echo.Context) error {
	all := "All bets:\n"
	bets, err := e.repository.GetAll()
	if err != nil {
		return err
	}
	for _, bet := range bets {
		all += bet.String()
	}
	return c.String(http.StatusOK, all)
}
