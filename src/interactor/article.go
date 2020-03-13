package interactor

import (
	"time"
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"
)

type ArticleInteractor struct {
	Repository                repository.IArticleRepository
	ArticleCategoryRepository repository.IArticleCategoryRepository
}

func NewArticleInteractor(
	repository repository.IArticleRepository,
	articleCategoryRepository repository.IArticleCategoryRepository,
) usecase.IArticleUsecase {
	return &ArticleInteractor{
		Repository:                repository,
		ArticleCategoryRepository: articleCategoryRepository,
	}
}

func (a *ArticleInteractor) Create(ra *request.CreateArticleRequest) uint64 {
	article := model.Article{
		Title:    ra.Title,
		Body:     ra.Body,
		Status:   ra.Status,
		UserID:   ra.UserID,
		ShopID:   ra.ShopID,
		CreateAt: time.Now(),
	}

	a.Repository.Save(&article)
	a.ArticleCategoryRepository.SaveCategoriesForArticle(article.ID, ra.CategoryIDs)

	return article.ID
}

func (a *ArticleInteractor) GetAll() []model.Article {
	return a.Repository.FindAll()
}
