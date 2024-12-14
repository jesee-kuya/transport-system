package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

const SCHOOL_TABLE = `
CREATE TABLE IF NOT EXISTS schools (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);
`

const PARENT_TABLE = `
CREATE TABLE IF NOT EXISTS parents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	full_name TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL,
	school TEXT NOT NULL,
	child_admission_number TEXT NOT NULL,
	password TEXT NOT NULL
);
`

var Tables = []string{SCHOOL_TABLE, PARENT_TABLE}

func (m *UserModel) InitTable() {
	for _, query := range Tables {
		_, err := m.DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to initialize tables: %v", err)
		}
	}
	log.Println("Tables initialized successfully")
}

func (m *UserModel) InsertSchool(name, email, password string) (bool, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("failed to hash password: %v", err)
	}

	_, err = m.DB.Exec(`INSERT INTO schools (name, email, password) VALUES (?, ?, ?)`, name, email, string(hashedPassword))
	if err != nil {
		return false, fmt.Errorf("failed to insert school: %v", err)
	}
	return true, nil
}

func (m *UserModel) InsertParent(fullName, email, school, childAdmissionNumber, password string) (bool, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("failed to hash password: %v", err)
	}

	_, err = m.DB.Exec(`INSERT INTO parents (full_name, email, school, child_admission_number, password) VALUES (?, ?, ?, ?, ?)`,
		fullName, email, school, childAdmissionNumber, string(hashedPassword))
	if err != nil {
		return false, fmt.Errorf("failed to insert parent: %v", err)
	}
	return true, nil
}

func (m *UserModel) CheckCredentials(email, password, userType string) (bool, error) {
	var hashedPassword string
	query := ""

	if userType == "school" {
		query = `SELECT password FROM schools WHERE email = ?`
	} else if userType == "parent" {
		query = `SELECT password FROM parents WHERE email = ?`
	} else {
		return false, fmt.Errorf("invalid user type")
	}

	err := m.DB.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No user found
		}
		return false, fmt.Errorf("failed to query credentials: %v", err)
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil // Password does not match
	}

	return true, nil // Credentials are valid
}
