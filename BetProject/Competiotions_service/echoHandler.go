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

func (e *EchoServer) Register() {
	//create setResult getAll getByID
	e.GET("create", e.create)
	e.GET("setResult", e.setResult)
	e.GET("getAll", e.getAll)
	e.GET("getById", e.getById)
	e.GET("delete", e.delete)
}
func (e *EchoServer) create(c echo.Context) error {
	comp := Competition{Name: c.QueryParam("name")}
	str, err := e.repository.Create(&comp)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, str)
}
func (e *EchoServer) setResult(c echo.Context) error {
	id := c.QueryParam("id")
	result := c.QueryParam("result")
	Id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	str, err := e.repository.SetResult(Id, result)
	if err != nil {
		return err
	}
	return c.String(http.StatusResetContent, str)
}
func (e *EchoServer) getAll(c echo.Context) error {
	all := "All competitions:\n"
	comps, err := e.repository.GetAll()
	if err != nil {
		return err
	}
	for _, comp := range comps {
		all += comp.String()
	}
	return c.String(http.StatusOK, all)
}
func (e *EchoServer) getById(c echo.Context) error {
	id := c.QueryParam("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	comp, err := e.repository.GetById(Id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, comp.String())
}
func (e *EchoServer) delete(c echo.Context) error {
	id := c.QueryParam("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	str, err := e.repository.Delete(Id)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, str)
}
