package server

import "auth_service/internal/database"

type Config struct {
	Port     string
	Database database.Config
}
