package handler

import (
	"time"
	"fmt"
	"torimo-article-api/src/model"
	// "os"
	"net/http"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo"
)

type RequestArticle struct {
	Title string           `json:"title"  validate:"required,max=100"`
	Body string            `json:"body"   validate:"required"`
	Status string          `json:"status" validate:"required,max=16"`
	UserID int64           `json:"userId" validate:"required"`
	ShopID string          `json:"shopId"`
	CategoryIDs []int      `json:"categoryIds"`
}

// CreateArticle return error
func (h *Handler) CreateArticle(c echo.Context) error {
	ra := new(RequestArticle)

	if err := c.Bind(ra); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(ra); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.(validator.ValidationErrors).Error())
	}

	// *** request file include. ***
	// avatar, err := c.FormFile("image")
  	// if err != nil {
 	// 	return err
	// }

	category := []model.Category{}
	h.db.Find(&category, "id IN (?)", ra.CategoryIDs)
	fmt.Println("category ids: ", ra.CategoryIDs)
	fmt.Println("result: ", category)

	shop := model.Shop{}
	h.db.Find(&shop, "id=?", ra.ShopID)

	user := model.User{}
	h.db.Find(&user, "id=?", ra.UserID)

	article := model.Article{
		Title: ra.Title,
		Body: ra.Body,
		Status: ra.Status,
		User: user,
		Shop: shop,
		Categories: category,
		CreateAt: time.Now(),
	}

	h.db.Create(&article)

	// TODO: avatar to AWS S3
	// fmt.Println(avatar)

	return c.JSON(http.StatusCreated, article)
}

// GetArticle return error
func (h *Handler) GetArticle(c echo.Context) error {
	article := model.Article{}
	h.db.Find(&article, "id=?", 1)

	// category := []model.Category{}
	// h.db.Find(&category, "id IN (?)", []int{1, 2})

	// shop := model.Shop{}
	// h.db.Find(&shop, "id=?", 1)

	// user := model.User{}
	// h.db.Find(&user, "id=?", 1)

	// a := model.Article{
	// 	Title: "Jinzhu",
	// 	Body: "string",
	// 	Status: "string",
	// 	User: user,
	// 	Shop: shop,
	// 	Categories: []*model.Category{&category},
	// 	CreateAt: time.Now(),
	// }

	// h.db.Create(&a)

	return c.JSON(http.StatusOK, article)
}
