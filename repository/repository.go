package repository

import "github.com/jmoiron/sqlx"

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &AuthRepositoryStruct{DB: db}
}

func NewSchoolRepository(db *sqlx.DB) SchoolRepository {
	return &SchoolRepositoryStruct{DB: db}
}

func NewGuardianRepository(db *sqlx.DB) GuardianRepository {
	return &GuardianRepositoryStruct{DB: db}
}

func NewPrivateParentRepository(db *sqlx.DB) PrivateParentRepository {
	return &PrivateParentRepositoryStruct{DB: db}
}

func NewSchoolDriverRepository(db *sqlx.DB) SchoolDriverRepository {
	return &SchoolDriverRepositoryStruct{DB: db}
}

func NewPrivateDriverRepository(db *sqlx.DB) PrivateDriverRepository {
	return &PrivateDriverRepositoryStruct{DB: db}
}
