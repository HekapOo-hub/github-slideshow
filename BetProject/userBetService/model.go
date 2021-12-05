package main

import "fmt"

type User struct {
	UserId   int
	Name     string
	Password string
	Balance  int
	Bets     string
	Role     string
}

func (u *User) String() string {
	return fmt.Sprintf("UserId:%d.UserName:%s. Balance:%d.\n Bets ids:%s\n", u.UserId, u.Name, u.Balance, u.Bets)
}

type SignInInfo struct {
	Name     string
	Password string
	Token    string
}
