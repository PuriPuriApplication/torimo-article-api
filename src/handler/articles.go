package handler

import (
	// "torimo-article-api/src/model"
	"net/http"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo"
)

type RequestArticle struct {
	Title string           `json:"title"  validate:"required,max=100"`
	Body string            `json:"body"   validate:"required"`
	Status string          `json:"status" validate:"required,max=16"`
	UserID int64           `json:"userId" validate:"required"`
	ShopName string        `json:"shopName"`
	CategoryNames []string `json:"categoryNames"`
}

// CreateArticle return error
func (h *Handler) CreateArticle(c echo.Context) error {
	// a := &article{
	// 	ID: seq,
	// }
	// if err := c.Bind(a); err != nil {
	// 	return err
	// }
	// articles[a.ID] = a
	// seq++

	ra := new(RequestArticle)

	if err := c.Bind(ra); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(ra); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}

	return c.JSON(http.StatusCreated, ra)
}

// GetArticle return error
// func (h *Handler) GetArticle(c echo.Context) error {
// 	a := model.Article{}
// 	h.db.Find(&a, "id=?", 1)

// 	return c.JSON(http.StatusOK, a)
// }
