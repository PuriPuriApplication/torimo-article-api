package handler

import (
	"net/http"
	"strconv"

	"torimo-article-api/src/domain/presenter"
	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

// Handler is Struct
type ArticleHandler struct {
	ArticleUsecase   usecase.IArticleUsecase
	ArticlePresenter presenter.IArticlePresenter
}

func NewArticleHandler(au usecase.IArticleUsecase, ap presenter.IArticlePresenter) *ArticleHandler {
	return &ArticleHandler{
		ArticleUsecase:   au,
		ArticlePresenter: ap,
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

	return c.JSON(http.StatusOK, ah.ArticlePresenter.MappingAll(&articles))
}

// GetOne return error
func (ah *ArticleHandler) GetOne(c echo.Context) error {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	article := ah.ArticleUsecase.GetOne(ID)

	return c.JSON(http.StatusOK, ah.ArticlePresenter.MappingOne(&article))
}
