package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error hashing %v\n", password)
		return nil, fmt.Errorf("Create user %w", err)
	}

	passwordHash := string(hashBytes)

	row, err := us.DB.Query(`
    INSERT INTO users (email, password)
    VALUES ($1, $2) RETURNING id`, email, password)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
