package main

import "fmt"

type Bet struct {
	CompetitionId int
	UserId        int
	Result        string
	Money         int
}

func (b *Bet) String() string {
	return fmt.Sprintf("CompetitionId: %d. UserId: %d.Money: %d.Result: %s\n", b.CompetitionId, b.UserId, b.Money, b.Result)
}
