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

	return c.String(http.StatusOK, message)
}

// Deposits funds into the balance
func depositHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	message, err := deposit(jsonMap)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, message)
}

func addUserHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	message, err := addUser(jsonMap)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, message)
}

func deleteHandler(c echo.Context) error {
	var jsonMap map[string]interface{}

	if err := c.Bind(&jsonMap); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := deleteUser(jsonMap); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

// The declaration of all routes comes from it
func routes(e *echo.Echo) {
	e.GET("/ping", pingHandler)
	e.POST("/addUser", addUserHandler)
	e.PATCH("/deposit", depositHandler)
	e.DELETE("/delete", deleteHandler)
}

func server() {
	e := echo.New()
	routes(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	log.Fatal(e.Start(":" + getEnvValue("PORT")))
}
