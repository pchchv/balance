package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

// Load values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

// Getting a value. Outputs a panic if the value is missing
func getEnvValue(v string) string {
	value, exist := os.LookupEnv(v)
	if !exist {
		log.Panicf("Value %v does not exist", v)
	}
	return value
}

// Deposits funds into the balance
func deposit(jsonMap map[string]interface{}) (map[string]string, error) {
	var balance, totalBalance, reserved float64
	id := fmt.Sprint(jsonMap["id"])
	funds, err := strconv.ParseFloat(fmt.Sprint(jsonMap["funds"]), 64)
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("select * from Users where id = $1", id)
	err = row.Scan(&id, &balance, &totalBalance, &reserved)
	if err != nil {
		return nil, err
	}

	balance += funds
	totalBalance += funds
	result := map[string]string{
		"id":      id,
		"balance": fmt.Sprint(balance),
	}

	_, err = db.Exec("update Users set (balance = $1, totalBalance = $2) where id = $3", balance, totalBalance, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Adds a user
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

	_, err = db.Exec("insert into Users (id, balance) values ($1, $2)",
		id, funds)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Deletes a user
func deleteUser(jsonMap map[string]interface{}) (map[string]string, error) {
	uuid := jsonMap["id"]
	if uuid == nil {
		return nil, fmt.Errorf("wrong id")
	}
	id := fmt.Sprint(uuid)

	_, err := db.Exec("delete from Users where id = $1", id)
	if err != nil {
		return nil, err
	}

	result := map[string]string{
		"id":      id,
		"deleted": "OK",
	}

	return result, nil
}

// Reserves funds
func reserve(jsonMap map[string]interface{}) (map[string]string, error) {
	var reserved, totalBalance float64
	balance := 111.11
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

	row := db.QueryRow("select * from Users where id = $1", user)
	err = row.Scan(&user, &balance, &totalBalance, &reserved)
	if err != nil {
		return nil, err
	}

	if balance < cost {
		return nil, fmt.Errorf("the amount on the balance is less than the cost of the service! balance: %v, cost: %v", balance, cost)
	}

	totalBalance = balance
	balance -= cost
	reserved += cost

	result := map[string]string{
		"id":            user,
		"balance":       fmt.Sprint(balance),
		"reserved":      fmt.Sprint(reserved),
		"total balance": fmt.Sprint(totalBalance),
	}

	_, err = db.Exec("update Users set (balance = $1, totalBalance = $2, reserved = $3) where id = $4", balance, totalBalance, reserved, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Charges the reserved funds from the user's account
func receipt(jsonMap map[string]interface{}) (map[string]string, error) {
	var balance, reserved, cost, totalBalance float64
	serviceID := jsonMap["serviceID"]
	if serviceID == nil {
		return nil, fmt.Errorf("wrong service id")
	}
	service := fmt.Sprint(serviceID)

	row := db.QueryRow("select * from Services where id = $1", serviceID)
	err := row.Scan(&service, &cost)
	if err != nil {
		return nil, err
	}

	uid := jsonMap["userID"]
	if uid == nil {
		return nil, fmt.Errorf("wrong user id")
	}
	user := fmt.Sprint(uid)

	row = db.QueryRow("select * from Users where id = $1", user)
	err = row.Scan(&user, &balance, &totalBalance, &reserved)
	if err != nil {
		return nil, err
	}

	totalBalance -= cost
	reserved -= cost

	result := map[string]string{
		"id":            user,
		"balance":       fmt.Sprint(balance),
		"reserved":      fmt.Sprint(reserved),
		"total balance": fmt.Sprint(totalBalance),
	}

	_, err = db.Exec("update Users set (balance = $1, totalBalance = $2, reserved = $3) where id = $4", balance, totalBalance, reserved, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Gets the user balance
func balance(jsonMap map[string]interface{}) (map[string]string, error) {
	var balance, totalBalance, reserved float64
	uuid := jsonMap["id"]
	if uuid == nil {
		return nil, fmt.Errorf("wrong id")
	}
	id := fmt.Sprint(uuid)

	row := db.QueryRow("select * from Users where id = $1", id)
	err := row.Scan(&id, &balance, &totalBalance, &reserved)
	if err != nil {
		return nil, err
	}

	result := map[string]string{
		"id":            id,
		"balance":       fmt.Sprint(balance),
		"reserved":      fmt.Sprint(reserved),
		"total balance": fmt.Sprint(totalBalance),
	}

	return result, nil
}

func main() {
	database()
	server()
}
