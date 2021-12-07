package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

var config = mysql.Config{
	User:      "hekapoo",
	Passwd:    "1234",
	Net:       "tcp",
	Addr:      "localhost:3306",
	DBName:    "competition",
	Collation: "",
}

type repository interface {
	CreateUser(*User) (string, error)
	GetByName(name string) (*User, error)
	GetAll() ([]*User, error)
	DeleteUser(id int) (string, error)
	SetBalance(id int, addMoney int) (string, error)
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
func (r *Repository) CreateUser(u *User) (string, error) {
	userStr := "user"
	empty := ""
	_, err := r.db.Exec("insert into users (name,password,balance,bets,role) values (?,?,?,?,?)", u.Name,
		u.Password, u.Balance, empty, userStr)
	if err != nil {
		return "", err
	}
	return "User was successfully created", nil
}
func (r Repository) GetAll() ([]*User, error) {
	rows, err := r.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*User, 0)
	for rows.Next() {
		u := User{}
		empty := make([]string, 2)
		err := rows.Scan(&u.UserId, &u.Name, &empty[0], &u.Balance, &u.Bets, &empty[1])
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
func (r *Repository) DeleteUser(id int) (string, error) {
	res, err := r.db.Exec("delete from users where id=?", id)
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
	return "user deleted", nil
}
func (r *Repository) GetByName(name string) (*User, error) {
	rows := r.db.QueryRow("SELECT * from users where name=?", name)
	u := User{}

	err := rows.Scan(&u.UserId, &u.Name, &u.Password, &u.Balance, &u.Bets, &u.Role)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
func (r *Repository) SetBalance(id int, addMoney int) (string, error) {
	_, err := r.db.Exec("update users set balance=balance+? where id=?", addMoney, id)
	if err != nil {
		return "database error", err
	}

	return "Balance was changed", nil
}
