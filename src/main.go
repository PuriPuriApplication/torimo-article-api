package main

import (
	"net/http"
	"torimo-article-api/src/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// create instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// Allows requests from any `http://localhost:8888` or `https://torimo-a04a5.firebaseapp.com` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8888", "https://torimo-a04a5.firebaseapp.com"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Handler
	h := &handler.Handler{}

	e.POST("/articles", h.CreateArticle)
	e.GET("/articles/:id", h.GetArticle)
	e.PUT("/articles/:id", h.UpdateArticle)
	e.DELETE("/articles/:id", h.DeleteArticle)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
