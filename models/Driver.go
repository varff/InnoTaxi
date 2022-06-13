package models

import (
	"InnoTaxi/connection"
	"context"
	"errors"
	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Driver struct {
	username string
	password string
	phone    int16
	email    string
	taxi     TaxiType
	Status   Status
	rate     int8
}
type TaxiType int64

const (
	Eco TaxiType = iota
	Comfort
	Business
)

func (s TaxiType) String() string {
	switch s {
	case Eco:
		return "Economy"
	case Comfort:
		return "Comfort"
	case Business:
		return "Business"
	}
	return "unknown"
}

type Status bool

const (
	busy Status = true
	free Status = false
)

func (s Status) String() string {
	switch s {
	case busy:
		return "Busy"
	case free:
		return "Free"
	}
	return "unknown"
}

func DriverRegistration(uname string, pass string, phone int16, email string, taxi TaxiType) error {
	db, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	defer db.Close()
	req := "insert into driver (username, password, phone, email, taxi, status, rate) values($1, $2, $3, $4,$5,$6,$7);"
	row := db.QueryRow(context.Background(), req, uname, pass, phone, email, taxi, free, 0)
	err := row.Scan()
	if err != nil {
		return err
	}
	return nil
}

func LoginDriver(uname string, pass string) error {
	var result Driver
	db, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	defer db.Close()
	req := "SELECT * FROM driver where username = $1;"
	row := db.QueryRow(context.Background(), req, uname)
	err := row.Scan(&result)
	if result.password != pass {
		return errors.New("wrong password")
	}
	return err //+token
}

func (driver *Driver) CheckRate() int8 {
	var rate int8
	db, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	defer db.Close()
	req := "SELECT rate FROM driver where username = $1;"
	row := db.QueryRow(context.Background(), req, driver.username)
	err := row.Scan(&rate)
	if err != nil {
		return 0
	}
	return rate
}

func (driver *Driver) FreeStatus() {
	dbDriver, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	dbOrder, _ := pgxpool.Connect(context.Background(), connection.OrderConString())
	defer dbDriver.Close()
	defer dbOrder.Close()
	reqDr := "UPDATE driver SET status='Free' where username=$1;"
	_, err1 := dbDriver.Exec(context.Background(), reqDr, driver.username)
	driver.Status = free
	reqOrder := "UPDATE order SET status='Free' where username=$1 AND status='Busy'"
	_, err2 := dbOrder.Exec(context.Background(), reqOrder, driver.username)
	if err2 != nil || err1 != nil {
		log.Fatal(err1, err2)
		return
	}
}

func (driver *Driver) RateOrderByDriver(order Order, rating int8) {
	//time
	order.driversRate = rating
}
