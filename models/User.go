package models

import (
	"InnoTaxi/connection"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type User struct {
	username  string
	Password  string
	phone     int32
	email     string
	rate      int32
	isAnalyst string
}

func UserLogin(phone int32, pass string) (bool, error) {
	var result User
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	que := "SELECT * FROM \"user\" where phone = $1;"
	row := db.QueryRow(context.Background(), que, phone)
	err := row.Scan(&result)
	if err == pgx.ErrNoRows {
		return false, nil
	}
	if err != nil {
		log.Printf("User Login %s", err)
		return false, err
	}
	if result.Password != pass {
		return false, nil
	}
	return true, nil
}

func UserRegister(Name, Password, Email string, Phone int32) (bool, error) {
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	que := "insert into \"user\" (\"name\", \"password\", phone, email, rate, analyst) " +
		"values($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(context.Background(), que, Name, Password, Phone, Email, 0, false)
	if err != nil {
		log.Printf("UserRegister %s", err)
		return false, err
	}
	return true, nil
}

func UserCheckRate(phone int) (int32, error) {
	var rate int32
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	que := "SELECT rate FROM user where phone = $1;"
	err := db.QueryRow(context.Background(), que, phone).Scan(&rate)
	if err != nil {
		log.Printf("UserCheckRate %s", err)
		return 0, err
	}
	return rate, nil
}

func UserRateOrder(phone, rate int32) error {
	return nil
}
