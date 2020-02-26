package interactor

import (
	"go.uber.org/dig"
	"time"

	"torimo-article-api/src/handler/request"
	"torimo-article-api/src/domain/repository"
	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/usecase"
)

type ArticleInteractor struct {
	dig.In
	Repository repository.IArticleRepository
}

func NewArticleInteractor(repository repository.IArticleRepository) usecase.IArticleUsecase {
	return &ArticleInteractor{
		Repository: repository,
	}
}

func (a *ArticleInteractor) Create(ra *request.RequestArticle) model.Article {
	article := model.Article{
		Title: ra.Title,
		Body: ra.Body,
		Status: ra.Status,
		User: model.User{
			Id: 1,
			Name: "1",
			ExternalServiceID: "1",
			ExternalServiceType: "1",
			IsDeleted: false,
			CreateAt: time.Now(),
			UpdateAt: nil,
			DeleteAt: nil,
		},
		Shop: model.Shop{
			ID: 1,
			Name: "1",
			StationID: 1,
			CreateUser: 1,
			IsDeleted: false,
			CreateAt: time.Now(),
			DeleteAt: nil,
		},
		Categories: []model.Category{
			{
				ID: 1,
				Name: "1",
				CreateUser: 1,
				IsDeleted: false,
				CreateAt: nil,
				DeleteAt: nil,
			},
		},
		CreateAt: time.Now(),
	}

	a.Repository.Save(&article)
	return article
}
