package main

import (
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

	// Handler
	h := &handler.Handler{}

	e.POST("/articles", h.CreateArticle)
	e.GET("/articles/:id", h.GetArticle)
	e.PUT("/articles/:id", h.UpdateArticle)
	e.DELETE("/articles/:id", h.DeleteArticle)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
