package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	CreatedAt    string `json:"created_at"`
}

func (u *User) Validate() error {
	if len(u.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}
	if !regexp.MustCompile(`(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`).MatchString(u.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(u.PasswordHash) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(bytes)
	return nil
}
