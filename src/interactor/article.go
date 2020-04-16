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

func (a *ArticleInteractor) Create(ra *request.UpsertArticleRequest) uint64 {
	article := model.Article{
		Title:    ra.Title,
		Body:     ra.Body,
		Status:   ra.Status,
		UserID:   ra.UserID,
		ShopID:   ra.ShopID,
		CreateAt: time.Now(),
	}

	a.Repository.Create(&article)
	a.ArticleCategoryRepository.SaveCategoriesForArticle(article.ID, ra.CategoryIDs)

	return article.ID
}

func (a *ArticleInteractor) Update(id uint64, ra *request.UpsertArticleRequest) {
	article := model.Article{
		ID:       id,
		Title:    ra.Title,
		Body:     ra.Body,
		Status:   ra.Status,
		UserID:   ra.UserID,
		ShopID:   ra.ShopID,
		CreateAt: time.Now(),
	}

	a.ArticleCategoryRepository.DeleteByAriticleID(article.ID)
	a.Repository.Update(&article)
	a.ArticleCategoryRepository.SaveCategoriesForArticle(article.ID, ra.CategoryIDs)
}

func (a *ArticleInteractor) GetAll() []model.Article {
	return a.Repository.FindAll()
}

func (a *ArticleInteractor) GetOne(ID uint64) model.Article {
	return a.Repository.FindOne(ID)
}
