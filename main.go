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
		return nil, fmt.Errorf("wrong id")
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
	var balance, reserved float64
	uid := jsonMap["userID"]
	if uid == nil {
		return nil, fmt.Errorf("wrong user id")
	}
	user := fmt.Sprint(uid)
	serviceID := jsonMap["serviceID"]
	if serviceID == nil {
		return nil, fmt.Errorf("wrong service id")
	}
	// service := fmt.Sprint(serv)
	orderID := jsonMap["orderID"]
	if orderID == nil {
		return nil, fmt.Errorf("wrong order id")
	}
	// order := fmt.Sprint(orderID)
	c := jsonMap["cost"]
	if c == nil {
		return nil, fmt.Errorf("wrong cost")
	}
	cost, err := strconv.ParseFloat(fmt.Sprint(c), 64)
	if err != nil {
		return nil, fmt.Errorf("wrong cost %v", err)
	}

	// TODO: Implement getting balance data from the database

	if balance < cost {
		return nil, fmt.Errorf("the amount on the balance is less than the cost of the service! balance: %v, cost: %v", balance, cost)
	}

	totalBalance := balance
	balance -= cost
	reserved += cost

	result := map[string]string{
		"id":            user,
		"balance":       fmt.Sprint(balance),
		"reserver":      fmt.Sprint(reserved),
		"total balance": fmt.Sprint(totalBalance),
	}

	// TODO: Implement data entry in the database

	return result, nil
}

func receipt(jsonMap map[string]interface{}) (map[string]string, error) {
	return nil, nil
}

func main() {
	server()
}
