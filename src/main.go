package main

import (
	"net/http"
	"torimo-article-api/src/handler"
	"torimo-article-api/src/validate"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"torimo-article-api/src/db"
)

// type CustomValidator struct {
// 	validator *validator.Validate
// }

// func (cv *CustomValidator) Validate(i interface{}) error {
// 	return cv.validator.Struct(i)
// }

func main() {

	// create instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// custom validator
	e.Validator = validate.NewValidator()
	// e.Validator = &CustomValidator{validator: validator.New()}

	// CORS restricted
	// Allows requests from any `http://localhost:8888` or `https://torimo-a04a5.firebaseapp.com` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8888", "https://torimo-a04a5.firebaseapp.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Handler
	d := db.Init()
	h := handler.NewHandler(d)

	e.POST("/articles", h.CreateArticle)
	// e.GET("/articles", h.GetArticle)
	// e.GET("/articles/:id", h.GetArticle)
	// e.PUT("/articles/:id", h.UpdateArticle)
	// e.DELETE("/articles/:id", h.DeleteArticle)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
