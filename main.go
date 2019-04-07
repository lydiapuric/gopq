package main

import (
	"fmt"
	"os"

	"./db"
)

func main() {

	// read configuration
	config, err := readConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// connect the datapase
	db, err := db.New(config)
	if err != nil {
		fmt.Printf("Could not connect to DB: %v\n", err)
		os.Exit(1)
	}

	// check if database can be reached
	if err := db.Ping(); err != nil {
		fmt.Printf("Could not reach the database %v\n", err)
		os.Exit(1)
	}

}
