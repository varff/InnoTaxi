package models

import (
	"InnoTaxi/connection"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	username string
	password string
	phone    int16
	email    string
	rate     int8
}

func RegisterUser(username string, pass string, phone int16, email string) error {
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	req := "insert into user (username, password, phone, email, rate) values($1, $2, $3, $4, $5);"
	_, err := db.Exec(context.Background(), req, username, pass, phone, email, 0)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(uname string, pass string) error {
	var result User
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	req := "SELECT * FROM user where username = $1;"
	row := db.QueryRow(context.Background(), req, uname)
	err := row.Scan(&result)
	if result.password != pass {
		return errors.New("wrong password")
	}
	return err
}

func (user *User) CheckRate() int8 {
	var rate int8
	db, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	defer db.Close()
	req := "SELECT rate FROM user where username = $1;"
	row := db.QueryRow(context.Background(), req, user.username)
	err := row.Scan(&rate)
	if err != nil {
		return 0
	}
	return rate
}

func (user *User) RateOrderByUser(order Order, rating int8) {
	//time
	order.usersRate = rating
}
