package repositories

import (
	user "authentication_service/genproto/authentication_service"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Login(userReq *user.LoginRequest) (*user.LoginResponse, error) {
	res := user.LoginResponse{}
	query :=
		`
			SELECT password_hash from users WHERE email = $1
		`

	log.Println(userReq.Password)
	var password_hash string
	err := r.db.QueryRow(query, userReq.Email).Scan(&password_hash)
	if err != nil {
		if err == sql.ErrNoRows {
			res.Success = false
			return &res, fmt.Errorf("error getting user from database: not rows found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(userReq.Password))
	if err != nil {
		res.Success = false
		return &res, fmt.Errorf("invalid password")
	}

	res.Success = true
	return &res, nil
}

func (r *UserRepository) GetProfileById(userReq *user.UserIdRequest) (*user.UserIdResponse, error) {
	query := `select username, email, password_hash from users where id = $1`
	row := r.db.QueryRow(query, userReq.Id)
	var res user.UserIdResponse
	res.Profile = &user.Profile{}
	
	err := row.Scan(&res.Profile.Name, &res.Profile.Email, &res.Profile.Password)
	if err != nil {
		log.Println(err)
		return &user.UserIdResponse{Profile: &user.Profile{}}, err
	}
	return &res, nil
}

func (r *UserRepository) Register(userReq *user.RegisterRequest) (*user.RegisterResponse, error) {
	query := `INSERT INTO users (id, username, password_hash, email)values($1, $2, $3, $4)`

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
