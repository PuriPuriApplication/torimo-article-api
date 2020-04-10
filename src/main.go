package main

import (
	"torimo-article-api/src/handler/custom"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func echoUse(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8888", "https://torimo-a04a5.firebaseapp.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// custom validator
	e.Validator = &custom.CustomValidator{
		Validator: validator.New(),
	}
}

func main() {
	e := echo.New()
	echoUse(e)
	routing(e)
	e.Logger.Fatal(e.Start(":8080"))
}
