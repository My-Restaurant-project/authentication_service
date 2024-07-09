package repositories

import (
	user "authentication_service/genproto/authentication_service"

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
	email := userReq.Email
	password := bcrypt.CompareHashAndPassword([]byte(userReq.Password), []byte(userReq.Password))

	query := `select password from users where email = $1`
	err := r.db.QueryRow(query, email, password).Scan(&password)
	if err != nil {
		return nil, err
	}

	return &user.LoginResponse{
		Success: true,
	}, nil

}

func (r *UserRepository) GetProfileById(userReq *user.UserIdRequest) (*user.UserIdResponse, error) {
	query := `select username, email, password from users where id = $1`
	row := r.db.QueryRow(query, userReq.Id)
	var res user.UserIdResponse

	err := row.Scan(&res.Profile.Name, &res.Profile.Email, &res.Profile.Password)
	if err != nil {
		return nil, err
	}
	return &res, nil
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
