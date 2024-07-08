package repositories

import (
	user "github.com/Projects/Restaurant_Reservation_System/authentication_service/genproto/authentication_service"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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
	query := `INSERT INTO users (id, username, password, email)values($1, $2, $3, $4)`

	password, err := HashPassword(userReq.Profile.Password)
	if err != nil {
		return nil, err
	}
	id := uuid.NewString()
	userReq.Profile.Id = id

	userReq.Profile.Password = string(password)

	_, err = r.db.Exec(query, userReq.Profile.Id, userReq.Profile.Name, userReq.Profile.Password, userReq.Profile.Email)
	if err != nil {
		return nil, err
	}
	profile := &user.Profile{
		Id:        userReq.Profile.Id,
		Name:      userReq.Profile.Name,
		Email:     userReq.Profile.Email,
		Password:  userReq.Profile.Password,
		CreatedAt: userReq.Profile.CreatedAt,
		UpdatedAt: userReq.Profile.UpdatedAt,
	}
	return &user.RegisterResponse{Profile: profile}, nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
