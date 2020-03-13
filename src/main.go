package main

import (
	"torimo-article-api/src/driver"
	"torimo-article-api/src/handler/custom"

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
	d := driver.Init()
	// defer db.Close(d)

	h := Initialize(d)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.POST("/articles", h.CreateArticle)
	e.GET("/articles", h.GetAll)
	// e.GET("/articles/:id", h.GetArticle)
	// e.PUT("/articles/:id", h.UpdateArticle)
	// e.DELETE("/articles/:id", h.DeleteArticle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
