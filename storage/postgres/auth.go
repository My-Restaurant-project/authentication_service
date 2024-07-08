package postgres

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	DB *sql.DB
}

func NewAuth(db *sql.DB) *Auth {
	return &Auth{
		DB: db,
	}
}

func (auth *Auth) Authenticate(username, password string) (bool, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

}
