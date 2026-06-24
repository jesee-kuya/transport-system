package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jesee-kuya/transport-system/domain"
)

func (r *AuthRepositoryStruct) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (username, email, password_hash, role)
		VALUES (:username, :email, :password_hash, :role)
		RETURNING id, username, email, role, is_active, created_at, updated_at
	`
	rows, err := r.DB.NamedQueryContext(ctx, query, user)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	defer rows.Close()

	var created domain.User
	if rows.Next() {
		if err := rows.StructScan(&created); err != nil {
			return nil, fmt.Errorf("scan created user: %w", err)
		}
	}
	return &created, nil
}

func (r *AuthRepositoryStruct) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, username, email, password_hash, role, is_active, created_at, updated_at FROM users WHERE email = $1`
	err := r.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return &user, nil
}

func (r *AuthRepositoryStruct) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, username, email, password_hash, role, is_active, created_at, updated_at FROM users WHERE username = $1`
	err := r.DB.GetContext(ctx, &user, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	return &user, nil
}
