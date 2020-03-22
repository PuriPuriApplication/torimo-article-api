package interactor

import (
	"time"
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/article_repository"
	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"
)

type ArticleInteractor struct {
	Repository                article_repository.IArticleRepository
	ArticleCategoryRepository article_repository.IArticleCategoryRepository
}

func NewArticleInteractor(
	repository article_repository.IArticleRepository,
	articleCategoryRepository article_repository.IArticleCategoryRepository,
) usecase.IArticleUsecase {
	return &ArticleInteractor{
		Repository:                repository,
		ArticleCategoryRepository: articleCategoryRepository,
	}
}

func (a *ArticleInteractor) Create(ra *request.CreateArticleRequest) uint64 {
	article := article_model.Article{
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

func (a *ArticleInteractor) GetAll() []article_model.Article {
	return a.Repository.FindAll()
}

func (a *ArticleInteractor) GetOne(ID uint64) article_model.Article {
	return a.Repository.FindOne(ID)
}
