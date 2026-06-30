package repository

import "github.com/jmoiron/sqlx"

type AuthRepositoryStruct struct {
	DB *sqlx.DB
}

type SchoolRepositoryStruct struct {
	DB *sqlx.DB
}

type GuardianRepositoryStruct struct {
	DB *sqlx.DB
}

type PrivateParentRepositoryStruct struct {
	DB *sqlx.DB
}

type SchoolDriverRepositoryStruct struct {
	DB *sqlx.DB
}

type PrivateDriverRepositoryStruct struct {
	DB *sqlx.DB
}
