package interactor

import (
	"time"

	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"
)

type ArticleInteractor struct {
	Repository         repository.IArticleRepository
	UserRepository     repository.IUserRepository
	ShopRepository     repository.IShopRepository
	CategoryRepository repository.ICategoryRepository
}

func NewArticleInteractor(
	repository repository.IArticleRepository,
	userRepository repository.IUserRepository,
	shopRepository repository.IShopRepository,
	categoryRepository repository.ICategoryRepository,
) usecase.IArticleUsecase {
	return &ArticleInteractor{
		Repository:         repository,
		UserRepository:     userRepository,
		ShopRepository:     shopRepository,
		CategoryRepository: categoryRepository,
	}
}

func (a *ArticleInteractor) Create(ra *request.RequestArticle) model.Article {
	categories := a.CategoryRepository.FindByIdIn(ra.CategoryIDs)

	shop := a.ShopRepository.FindById(ra.ShopID)

	user := a.UserRepository.FindById(ra.UserID)

	article := model.Article{
		Title:      ra.Title,
		Body:       ra.Body,
		Status:     ra.Status,
		User:       user,
		Shop:       shop,
		Categories: categories,
		CreateAt:   time.Now(),
	}

	a.Repository.Save(&article)

	return article
}
