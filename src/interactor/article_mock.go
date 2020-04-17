package interactor

import (
	"fmt"
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/usecase"
)

type ArticleMockInteractor struct{}

func NewArticleMockInteractor() usecase.IArticleUsecase {
	return &ArticleMockInteractor{}
}

func (ami *ArticleMockInteractor) Create(ra *request.CreateArticleRequest) uint64 {
	fmt.Println("Create article")
	return 1
}

func (a *ArticleInteractor) Update(id uint64, ra *request.UpsertArticleRequest) {
	fmt.Println("Update article")
}

func (ami *ArticleMockInteractor) GetAll() []model.Article {
	return []model.Article{}
}

func (ami *ArticleMockInteractor) GetOne(ID uint64) model.Article {
	return model.Article{}
}
