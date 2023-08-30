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

	user := User{
		Email:        email,
		PasswordHash: password,
	}

	row := us.DB.QueryRow(`
    INSERT INTO users (email, password_hash)
    VALUES ($1, $2) RETURNING id`, email, passwordHash)
	if err != nil {
		return nil, err
	}

	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("Scaning user %w", err)
	}

	return &user, nil
}