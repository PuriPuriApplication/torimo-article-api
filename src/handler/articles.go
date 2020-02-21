package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	article struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}
)

var (
	articles = map[int]*article{}
	seq      = 1
)

func (h *Handler) createArticle(c echo.Context) error {
	a := &article{
		ID: seq,
	}
	if err := c.Bind(a); err != nil {
		return err
	}
	articles[a.ID] = a
	seq++
	return c.JSON(http.StatusCreated, a)
}

func (h *Handler) getArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, articles[id])
}

func (h *Handler) updateArticle(c echo.Context) error {
	a := new(article)
	if err := c.Bind(a); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	articles[id].Content = a.Content
	return c.JSON(http.StatusOK, articles[id])
}

func (h *Handler) deleteArticle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(articles, id)
	return c.NoContent(http.StatusNoContent)
}
