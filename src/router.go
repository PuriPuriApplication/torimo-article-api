package main

import (
	"net/http"
	"torimo-article-api/src/driver"

	"github.com/labstack/echo"
)

const (
	healthPath  string = "/health"
	articlePath string = "/articles"
)

func routing(e *echo.Echo) {
	// health
	e.GET(healthPath, func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// articles
	h := InitArticleHandler(driver.Init())
	e.POST(articlePath, h.CreateArticle)
	e.GET(articlePath, h.GetAll)
	e.GET(articlePath+"/:id", h.GetOne)
	// e.GET("/articles/:id", h.GetArticle)
	e.PUT("/articles/:id", h.UpdateArticle)
	// e.DELETE("/articles/:id", h.DeleteArticle)
}