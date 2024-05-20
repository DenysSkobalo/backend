package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	ID                int64        `json:"id"`
	Username          string       `json:"username"`
	Email             string       `json:"email"`
	Password          string       `json:"password"`
	ConfirmedPassword string       `json:"confirmed_password"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	Roles             []Role       `json:"roles"`
	Permissions       []Permission `json:"permissions"`
	CreatedAt         string       `json:"created_at"`
}

func (u *User) Validate() error {
	if len(u.Username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long")
	}
	if !regexp.MustCompile(`(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`).MatchString(u.Email) {
		return fmt.Errorf("invalid email format")
	}
	if len(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
