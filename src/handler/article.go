package handler

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"
)

// Handler is Struct
type ArticleHandler struct {
	dig.In
	ArticleUsecase  usecase.IArticleUsecase
	ResponseUsecase usecase.IResponseUsecase
}

func NewArticleHandler(au usecase.IArticleUsecase, ru usecase.IResponseUsecase) *ArticleHandler {
	return &ArticleHandler{
		ArticleUsecase:  au,
		ResponseUsecase: ru,
	}
}

// CreateArticle return error
func (ah *ArticleHandler) CreateArticle(c echo.Context) error {
	ra := new(request.CreateArticleRequest)

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

// GetAll return error
func (ah *ArticleHandler) GetAll(c echo.Context) error {
	articles := ah.ArticleUsecase.GetAll()

	return c.JSON(http.StatusOK, ah.ResponseUsecase.MappingAll(articles))
}

// GetOne return error
func (ah *ArticleHandler) GetOne(c echo.Context) error {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	article := ah.ArticleUsecase.GetOne(ID)

	return c.JSON(http.StatusOK, ah.ResponseUsecase.MappingOne(&article))
}
