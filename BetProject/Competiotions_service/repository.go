package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type repository interface {
	Create(*Competition) (string, error)
	SetResult(int, string) (string, error)
	GetAll() ([]*Competition, error)
	GetById(int) (*Competition, error)
	Delete(int) (string, error)
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
func (r *Repository) Create(c *Competition) (string, error) {
	var result string
	if c.Result == "" {
		result = "unknown"
	} else {
		result = c.Result
	}
	_, err := r.db.Exec("insert into Competition (name,result) values (?,?)", c.Name, result)
	if err != nil {
		return "", err
	}
	return "Competition successfully created!", nil
}
func (r *Repository) SetResult(id int, result string) (string, error) {
	res, err := r.db.Exec("update Competition set result =? where id=?", result, id)
	if err != nil {
		return "Database error", err
	}
	rowsAff, err := res.RowsAffected()
	if err != nil {
		return "Database error", err
	}
	if rowsAff == 0 {
		return "there is no such competition ", nil
	}
	return "result was set successfully", nil
}
func (r *Repository) GetAll() ([]*Competition, error) {
	rows, err := r.db.Query("select * from Competition")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comps := make([]*Competition, 0)
	for rows.Next() {
		c := Competition{}
		err := rows.Scan(&c.Id, &c.Name, &c.Result)
		if err != nil {
			return nil, err
		}
		comps = append(comps, &c)
	}
	return comps, nil
}
func (r *Repository) GetById(id int) (*Competition, error) {
	row := r.db.QueryRow("select * from Competition where id=?", id)
	c := Competition{}
	err := row.Scan(&c.Id, &c.Name, &c.Result)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *Repository) Delete(id int) (string, error) {
	res, err := r.db.Exec("delete from Competition where id=?", id)
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
	return "competition deleted", nil
}
