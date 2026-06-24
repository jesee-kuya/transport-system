package repository

import "github.com/jmoiron/sqlx"

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &AuthRepositoryStruct{DB: db}
}
