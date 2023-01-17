package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Checks that the server is up and running
func pingHandler(c echo.Context) error {
	message := "Balance service. Version 0.0.1"

	return c.JSON(http.StatusOK, message)
}

// Deposits funds into the balance
func depositHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	message, err := deposit(jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, message)
}

// Adds a user
func addUserHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	message, err := addUser(jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, message)
}

// Deletes a user
func deleteHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	message, err := deleteUser(jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, message)
}

// Reserves funds
func reserveHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	message, err := reserve(jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, message)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/ping", pingHandler)
	e.POST("/user", addUserHandler)
	e.PATCH("/deposit", depositHandler)
	e.DELETE("/user", deleteHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
