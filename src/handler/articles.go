package handler

import (
	"torimo-article-api/src/model"
	"net/http"

	"github.com/labstack/echo"
)

// CreateArticle return error
// func (h *Handler) CreateArticle(c echo.Context) error {
// 	a := &article{
// 		ID: seq,
// 	}
// 	if err := c.Bind(a); err != nil {
// 		return err
// 	}
// 	articles[a.ID] = a
// 	seq++
// 	return c.JSON(http.StatusCreated, a)
// }

// GetArticle return error
func (h *Handler) GetArticle(c echo.Context) error {
	a := model.Article{}
	h.db.Find(&a, "id=?", 1)

	return c.JSON(http.StatusOK, a)
}
