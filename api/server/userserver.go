package server

import (
	"InnoTaxi/api/generated/api"
	"InnoTaxi/connection"
	"InnoTaxi/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GRPCServer struct{}

func (s *GRPCServer) UserLogin(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	var result models.User
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	que := "SELECT * FROM user where phone = $1;"
	err := db.QueryRow(context.Background(), que, req.Phone).Scan(&result)
	if err != nil {
		return &api.LoginResponse{Token: ""}, err
	}
	if result.Password != req.Password {
		return &api.LoginResponse{Token: ""}, errors.New("wrong password")
	}
	return &api.LoginResponse{Token: ""}, nil
}

func (s *GRPCServer) UserRegister(ctx context.Context, req *api.RegisterRequest) (*api.LoginResponse, error) {
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	que := "insert into user (username, password, phone, email, rate) values($1, $2, $3, $4, $5);"
	_, err := db.Exec(context.Background(), que, req.Name, req.Password, req.Phone, req.Email, 0)
	if err != nil {
		return &api.LoginResponse{Token: ""}, err
	}
	return &api.LoginResponse{Token: ""}, nil
}

func (s *GRPCServer) UserCheckRate(ctx context.Context, req *api.Phone) (api.Rating, error) {
	var rate int32
	db, _ := pgxpool.Connect(context.Background(), connection.UserConString())
	defer db.Close()
	que := "SELECT rate FROM user where phone = $1;"
	err := db.QueryRow(context.Background(), que, req.Phone).Scan(&rate)
	if err != nil {
		return api.Rating{Rate: 0}, err
	}
	return api.Rating{Rate: rate}, nil
}
