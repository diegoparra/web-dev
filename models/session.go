package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when crating a new session
	// this will be left empty, as we onlt store the hash of a session Token
	// in our database and wen can't reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// TODO: Create a session token
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}
