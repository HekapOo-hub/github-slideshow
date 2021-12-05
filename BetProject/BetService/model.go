package main

import "fmt"

type Bet struct {
	CompetitionId int
	UserId        int
	Result        string
	Money         int
}

func (b *Bet) String() string {
	return fmt.Sprintf("CompetitionId: %d. UserId: %d. Result: %s.Money:%d\n", b.CompetitionId, b.UserId, b.Result, b.Money)
}
