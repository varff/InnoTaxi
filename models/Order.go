package models

import (
	"InnoTaxi/connection"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Order struct {
	user        string
	driver      string
	from, to    string
	taxiType    TaxiType
	status      OrderStatus
	time        time.Time
	usersRate   int8
	driversRate int8
}

func NewOrder(user string, driver string, from string, to string, taxiType TaxiType, time time.Time) error {
	dbDriver, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	dbOrder, _ := pgxpool.Connect(context.Background(), connection.DriverConString())
	defer dbDriver.Close()
	defer dbOrder.Close()

	reqDr := "UPDATE driver SET status='Busy' where username=$1;"
	_, err1 := dbDriver.Exec(context.Background(), reqDr, driver)
	if err1 != nil {
		return err1
	}

	reqOrder := "INSERT INTO order (user, driver, from, to, taxi, status, time) values($1, $2, $3, $4, $5, $6, $7, -1, -1)"
	_, err2 := dbOrder.Exec(context.Background(), reqOrder, user, driver, from, to, taxiType, inProgress, time)
	if err2 != nil {
		return err2
	}
	return nil
}

type OrderStatus bool

const (
	inProgress OrderStatus = true
	finished   OrderStatus = false
)

func (s OrderStatus) String() string {
	switch s {
	case inProgress:
		return "In Progress"
	case finished:
		return "Finished"
	}
	return "unknown"
}
