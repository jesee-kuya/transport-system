package repository

import "github.com/jmoiron/sqlx"

type AuthRepositoryStruct struct {
	DB *sqlx.DB
}
