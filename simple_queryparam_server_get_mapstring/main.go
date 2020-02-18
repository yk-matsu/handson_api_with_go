package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", func(c echo.Context) (err error) {
		names := c.ParamNames()
		values := c.ParamValues()
		params := map[string][]string{}
		for i, name := range names {
			params[name] = []string{values[i]}
		}
		return c.JSON(http.StatusOK, params)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
