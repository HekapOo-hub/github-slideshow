package main

import (
	"fmt"
)

type Competition struct {
	Id     int
	Name   string
	Result string
}

func (c *Competition) String() string {
	return fmt.Sprintf("CompetitionId: %d.Name: %s.Result: %s\n", c.Id, c.Name, c.Result)
}
