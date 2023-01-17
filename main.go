package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
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

func deposit(jsonMap map[string]interface{}) (map[string]string, error) {
	var balance float64
	id := fmt.Sprint(jsonMap["id"])
	funds, err := strconv.ParseFloat(fmt.Sprint(jsonMap["funds"]), 64)
	if err != nil {
		return nil, err
	}

	// TODO: Implement getting balance data from the database

	balance += funds
	result := map[string]string{
		"id":      id,
		"balance": fmt.Sprint(balance),
	}

	// TODO: Implement data entry in the database

	return result, nil
}

func addUser(jsonMap map[string]interface{}) (map[string]string, error) {
	f := jsonMap["funds"]
	var err error
	funds := 0.0

	if f != nil {
		funds, err = strconv.ParseFloat(fmt.Sprint(f), 64)
		if err != nil {
			return nil, err
		}
	}

	id := uuid.New()
	result := map[string]string{
		"id":      id.String(),
		"balance": fmt.Sprint(funds),
	}

	// TODO: Implement data entry in the database

	return result, nil
}

func deleteUser(jsonMap map[string]interface{}) (map[string]string, error) {
	uuid := jsonMap["id"]
	if uuid == nil {
		return nil, fmt.Errorf("Wrong ID!")
	}
	id := fmt.Sprint(uuid)

	// TODO: Implement deletion of user data from the database

	result := map[string]string{
		"id":      id,
		"deleted": "OK",
	}

	return result, nil
}

func reserve(jsonMap map[string]interface{}) (map[string]string, error) {
	user := jsonMap["userID"]
	if user == nil {
		return nil, fmt.Errorf("Wrong user ID!")
	}
	service := jsonMap["serviceID"]
	if service == nil {
		return nil, fmt.Errorf("Wrong service ID!")
	}
	order := jsonMap["orderID"]
	if order == nil {
		return nil, fmt.Errorf("Wrong order ID!")
	}
	cost := jsonMap["cost"]
	if cost == nil {
		return nil, fmt.Errorf("Wrong cost!")
	}
	return nil, nil
}

func main() {
	server()
}
