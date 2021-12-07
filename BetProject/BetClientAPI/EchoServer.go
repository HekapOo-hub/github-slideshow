package main

import (
	pb2 "BetClientAPI/Bet"
	pb1 "BetClientAPI/competition"
	pb3 "BetClientAPI/user"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type EchoServer struct {
	*echo.Echo
	pb1.CompetitionServiceClient
	pb2.BetServiceClient
	pb3.UserServiceClient
}

func (e *EchoServer) Register() {
	e.GET("/createUser", e.createUser)
	e.GET("/signIn", e.signIn)
	e.GET("/getAllUsers", e.getAllUsers)
	e.GET("/deleteUser", e.deleteUser)
	e.GET("/competition/create", e.createCompetition)
	e.GET("/competition/get", e.getCompetitionById)
	e.GET("/competition/getAll", e.getAllCompetitions)
	e.GET("/competition/setResult", e.setResult)
	e.GET("/bet/make", e.makeBet)
	e.GET("/bet/remove", e.removeBet)
	e.GET("/bet/getAll", e.getAllBets)
	e.GET("/bet/getMyBets", e.getMyBets)
	e.GET("/bet/checkResults", e.checkResults)
}
func (e *EchoServer) createUser(c echo.Context) error {
	u := &pb3.User{Name: c.QueryParam("name"), Password: c.QueryParam("password"), Balance: 0}
	res, err := e.UserServiceClient.CreateUser(context.Background(), u)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, res.String())
}
func (e *EchoServer) signIn(c echo.Context) error {
	name := c.QueryParam("name")
	pwd := c.QueryParam("password")
	res, err := e.UserServiceClient.SignIn(context.Background(), &pb3.SignInInfo{Name: name, Password: pwd})
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "token", Value: "", HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
		return c.String(http.StatusBadRequest, "There is no such user")
	}
	c.SetCookie(&http.Cookie{Name: "token", Value: res.Token, HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
	return c.String(http.StatusAccepted, "You entered successfully")
}
func (e *EchoServer) getAllUsers(c echo.Context) error {
	res, err := e.UserServiceClient.GetAll(context.Background(), &pb3.Empty{})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) deleteUser(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	res, err := e.UserServiceClient.DeleteUser(context.Background(), &pb3.Id{Id: int32(ID)})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) createCompetition(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value

	res, err := e.CompetitionServiceClient.CreateCompetition(context.Background(), &pb1.Competition{
		Name: c.QueryParam("name"), Result: c.QueryParam("result"), Token: token,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.String(http.StatusCreated, res.Response)
}
func (e *EchoServer) getCompetitionById(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	res, err := e.CompetitionServiceClient.GetById(context.Background(), &pb1.Id{
		Id: int32(ID), Token: token,
	})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) getAllCompetitions(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	res, err := e.CompetitionServiceClient.GetAll(context.Background(), &pb1.Empty{Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) setResult(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	res, err := e.CompetitionServiceClient.SetResult(context.Background(), &pb1.Request{
		Result: c.QueryParam("result"), Id: int32(ID), Token: token,
	})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) makeBet(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value

	cID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	money, err := strconv.Atoi(c.QueryParam("money"))
	if err != nil {
		return err
	}
	res, err := e.MakeBet(context.Background(), &pb2.BetWithToken{Bet: &pb2.Bet{
		CompetitionId: int32(cID), UserId: -1, Result: c.QueryParam("result"), Money: int32(money),
	}, Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, res.Response)
}
func (e *EchoServer) removeBet(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	cId, err := strconv.Atoi(c.QueryParam("cId"))
	uId, err := strconv.Atoi(c.QueryParam("uId"))
	if err != nil {
		return err
	}
	res, err := e.BetServiceClient.RemoveBet(context.Background(), &pb2.IdWithToken{CId: int32(cId),
		UId: int32(uId), Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) getAllBets(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	res, err := e.BetServiceClient.GetAll(context.Background(), &pb2.Empty{Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) getMyBets(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	res, err := e.BetServiceClient.GetMyBets(context.Background(), &pb2.IdWithToken{Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) checkResults(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	res, err := e.BetServiceClient.GetAll(context.Background(), &pb2.Empty{Token: token})
	if err != nil {
		return err
	}
	strBets := strings.Split(res.Response, "\n")
	fmt.Println(strBets[1])
	bets := make([]*pb2.Bet, 0)
	for i := 1; i < len(strBets)-1; i++ {
		var bet pb2.Bet
		str, err := fmt.Sscanf(strBets[i], "CompetitionId: %d. UserId: %d.Money: %d.Result: %s",
			&bet.CompetitionId, &bet.UserId, &bet.Money, &bet.Result)
		fmt.Println(str)
		fmt.Println(err)
		bets = append(bets, &bet)
	}
	for _, bet := range bets {
		res, err := e.CompetitionServiceClient.GetById(context.Background(), &pb1.Id{
			Id: bet.CompetitionId, Token: token})
		if err != nil {
			return err
		}
		var comp pb1.Competition
		fmt.Sscanf(res.Response, "CompetitionId: %d.Name: %s.Result: %s",
			comp.Id, comp.Name, comp.Result)
		if comp.Result != "unknown" {
			if comp.Result == bet.Result {
				_, err := e.UserServiceClient.SetBalance(context.Background(), &pb3.SetBalanceInfo{
					Id: bet.UserId, Money: bet.Money,
				})
				if err != nil {
					return err
				}
			} else {
				_, err := e.UserServiceClient.SetBalance(context.Background(), &pb3.SetBalanceInfo{
					Id: bet.UserId, Money: -1 * bet.Money,
				})
				if err != nil {
					return nil
				}
			}
		}
	}
	return c.String(http.StatusOK, "You checked results")
}
