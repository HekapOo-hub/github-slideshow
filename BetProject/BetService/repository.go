package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type repository interface {
	MakeBet(*Bet) (string, error)
	RemoveBet(int, int) (string, error)
	GetAll() ([]*Bet, error)
	GetMyBets(int) ([]*Bet, error)
}

var config = mysql.Config{
	User:      "hekapoo",
	Passwd:    "1234",
	Net:       "tcp",
	Addr:      "localhost:3306",
	DBName:    "competition",
	Collation: "",
}

type Repository struct {
	db *sql.DB
}

func NewRepository() repository {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil
	}
	return &Repository{db: db}
}
func (r *Repository) MakeBet(bet *Bet) (string, error) {
	_, err := r.db.Exec("insert into bets (competitionId,userId,result,money) values (?,?,?,?)",
		bet.CompetitionId, bet.UserId, bet.Result, bet.Money)
	if err != nil {
		return "", err
	}
	return "Your bet is made successfully", nil
}
func (r *Repository) RemoveBet(cId, uId int) (string, error) {
	res, err := r.db.Exec("delete from bets where competitionId=? && userId=?", cId, uId)
	if err != nil {
		return "", err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rows == 0 {
		return "there is no such competition", nil
	}
	return "bet deleted", nil
}
func (r *Repository) GetAll() ([]*Bet, error) {
	rows, err := r.db.Query("select * from bets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bets := make([]*Bet, 0)
	for rows.Next() {
		b := Bet{}
		err := rows.Scan(&b.CompetitionId, &b.UserId, &b.Result, &b.Money)
		if err != nil {
			return nil, err
		}
		bets = append(bets, &b)
	}
	return bets, nil
}
func (r *Repository) GetMyBets(id int) ([]*Bet, error) {
	rows, err := r.db.Query("select * from bets where userId=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bets := make([]*Bet, 0)
	fmt.Println(id)
	for rows.Next() {
		b := Bet{}
		var id int
		err := rows.Scan(&b.CompetitionId, &id, &b.Result, &b.Money)
		if err != nil {
			return nil, err
		}
		bets = append(bets, &b)
	}
	fmt.Println(bets)
	return bets, nil
}
