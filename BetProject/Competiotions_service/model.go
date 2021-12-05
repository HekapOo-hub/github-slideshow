package main

import "strconv"

type Competition struct {
	Id     int
	Name   string
	Result string
}

func (c *Competition) String() string {
	return "Competition " + strconv.Itoa(c.Id) + ": " + c.Name + ". Result: " + c.Result + "\n"
}
