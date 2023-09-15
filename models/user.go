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
		PasswordHash: passwordHash,
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

// Create the func Authenticate

func (us *UserService) Authenticate(email, password string) (*User, error) {
	// Make sure the email is all lowercase
	email = strings.ToLower(email)

	// Create an object of type User and assign the email that came from func variable
	user := User{
		Email: email,
	}

	// run the QueryRow to get the ID and PasswordHash from DB
	row := us.DB.QueryRow(`SELECT id, password_hash FROM users WHERE email=$1`, email)

	// run the result scan and assign the values to object User
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		fmt.Printf("Error on Scanning the user: %s", err)
		return nil, fmt.Errorf("Error getting information from DB: %w", err)
	}

	// Check the bcrypt compare hash and password to validade if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("Error comparing the password, %w", err)
	}

	// return the user and error
	return &user, nil
}
