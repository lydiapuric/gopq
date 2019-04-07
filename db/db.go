package db

import (
	"database/sql"
	"fmt"

	// enable the PostgreSQL driver
	_ "github.com/lib/pq"
)

// Settings for database
type Settings struct {
	Host   string
	Port   string
	DBName string
	User   string
	Pwd    string
	SSL    string
}

// New connection to PostgreSQL database
func New(s *Settings) (*sql.DB, error) {
	if s == nil {
		return nil, fmt.Errorf("Configuration is incomplete")
	}
	conn := "host=%s port=%s dbname=%s user=%s password=%s sslmode=%s"
	return sql.Open("postgres", fmt.Sprintf(conn, s.Host, s.Port, s.DBName, s.User, s.Pwd, s.SSL))
}
