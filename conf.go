package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./db"
)

// ConfigurationError - custom error for configuration
type ConfigurationError struct {
	errors []error
}

func (c *ConfigurationError) Error() string {
	messages := []string{}
	for _, err := range c.errors {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "\n")
}

// declare command-line flags to set the database configuration
var (
	dbHost = flag.String("h", "", "Host")
	dbPort = flag.String("p", "", "Port")
	dbName = flag.String("db", "", "Database")
	dbUser = flag.String("u", "", "Login Name")
	dbPwd  = flag.String("pwd", "", "Password")
	dbSSL  = flag.String("ssl", "", "SSL Mode (default: require)")
)

// read configuration for PostgreSQL database
func readConfiguration() (*db.Settings, error) {
	// parse the provided command-line flags
	flag.Parse()

	errors := []error{}

	// host
	host, found := getOption(dbHost, "DB_HOST")
	if !found {
		errors = append(errors, fmt.Errorf("Please specify host by using commandline flag -h or set it as environment variable DB_HOST"))
	}

	// port
	port, found := getOption(dbPort, "DB_PORT")
	if !found {
		errors = append(errors, fmt.Errorf("Please specify port by using commandline flag -p or set it as environment variable DB_PORT"))
	}

	// database name
	database, found := getOption(dbName, "DB_NAME")
	if !found {
		errors = append(errors, fmt.Errorf("Please specify database name by using commandline flag -db or set it as environment variable DB_NAME"))
	}

	// user
	user, found := getOption(dbUser, "DB_USER")
	if !found {
		errors = append(errors, fmt.Errorf("Please specify user by using commandline flag -u or set it as environment variable DB_USER"))
	}

	// password
	pwd, found := getOption(dbPwd, "DB_PWD")
	if !found {
		errors = append(errors, fmt.Errorf("Please specify password by using commandline flag -pwd or set it as environment variable DB_PWD"))
	}

	// ssl
	ssl, found := getOption(dbSSL, "DB_SSL")
	if !found {
		// set default ssl mode to require
		ssl = "require"
	}

	// configuration is not successful
	if len(errors) > 0 {
		return nil, &ConfigurationError{errors}
	}

	return &db.Settings{
		Host:   host,
		Port:   port,
		DBName: database,
		User:   user,
		Pwd:    pwd,
		SSL:    ssl,
	}, nil
}

// Get option
// If command-line flag is not set the fallback
// is to check for an environment variable
func getOption(flag *string, env string) (string, bool) {
	if *flag != "" {
		return *flag, true
	}
	if value := os.Getenv(env); value != "" {
		return value, true
	}
	return "", false
}
