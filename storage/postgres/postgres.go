package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Projects/Restaurant_Reservation_System/authentication_service/configs"
)

func ConnectDB() (*sql.DB, error) {
	dns := fmt.Sprintf("host=%s post=%s user=%s password=%s dbname=%s sslmode=disable", configs.Config{}.Postgres.DbHost)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
