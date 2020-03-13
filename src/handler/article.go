package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"go.uber.org/dig"

	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"
)

// Handler is Struct
type ArticleHandler struct {
	dig.In
	ArticleUsecase usecase.IArticleUsecase
}

func NewArticleHandler(au usecase.IArticleUsecase) *ArticleHandler {
	return &ArticleHandler{
		ArticleUsecase: au,
	}
}

// CreateArticle return error
func (ah *ArticleHandler) CreateArticle(c echo.Context) error {
	ra := new(request.RequestArticle)

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

	articleID := ah.ArticleUsecase.Create(ra)

	return c.JSON(http.StatusCreated, articleID)
}
