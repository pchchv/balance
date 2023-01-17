package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) string {
	// Getting a value. Outputs a panic if the value is missing
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

func deposit(jsonMap map[string]interface{}) error {
	id := fmt.Sprint(jsonMap["id"])
	funds, err := strconv.ParseFloat(fmt.Sprint(jsonMap["funds"]), 64)
	if err != nil {
		return err
	}
	// TODO: Implement data entry in the database

	return nil
}

func addUser(jsonMap map[string]interface{}) (string, error) {
	return "", nil
}

func main() {
	server()
}
