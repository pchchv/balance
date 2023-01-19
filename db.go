package main

import (
	"database/sql"
	"fmt"
)

func db() {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", getEnvValue("DBUSER"), getEnvValue("DBPASS"), getEnvValue("DBNAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
