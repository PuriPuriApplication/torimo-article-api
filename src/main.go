package main

import (
	"torimo-article-api/src/handler/custom"
	"torimo-article-api/src/infrastructure/db"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// create instance
	e := echo.New()

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

	// Handler
	d := db.Init()
	// defer db.Close(d)

	h := Initialize(d)

	// Init(d, e)

	e.POST("/articles", h.CreateArticle)
	// e.GET("/articles", h.GetArticle)
	// e.GET("/articles/:id", h.GetArticle)
	// e.PUT("/articles/:id", h.UpdateArticle)
	// e.DELETE("/articles/:id", h.DeleteArticle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
