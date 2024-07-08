package repositories

import (
	user "github.com/Projects/Restaurant_Reservation_System/authentication_service/genproto/authentication_service"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Login(userReq *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, nil
}

func (r *UserRepository) GetProfileById(userReq *user.UserIdRequest) (*user.UserIdResponse, error) {
	return nil, nil
}

func (r *UserRepository) Register(userReq *user.RegisterRequest) (*user.RegisterResponse, error) {
	return nil, nil
}
