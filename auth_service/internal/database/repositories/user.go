package repositories

import (
	"auth_service/pkg/models"
	"database/sql"
	"log"
)

type Repositories struct {
	UserRepo *UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepo: NewUserRepository(db),
	}
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(u *models.User) error {
	if err := u.HashPassword(u.PasswordHash); err != nil {
		log.Printf("Unable to hash password: %v", err)
		return err
	}

	const query = `INSERT INTO users (username, email, password_hash, created_at, first_name, last_name) VALUES ($1, $2, $3, NOW(), $4, $5)`
	_, err := repo.DB.Exec(query, u.Username, u.Email, u.PasswordHash, u.FirstName, u.LastName)
	if err != nil {
		log.Printf("Unable to create user: %v", err)
		return err
	}
	return nil
}
