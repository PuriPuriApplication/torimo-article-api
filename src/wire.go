//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/handler"
	"torimo-article-api/src/infrastructure/datastore"
	"torimo-article-api/src/interactor"
	"torimo-article-api/src/presenter"
)

func InitArticleHandler(d *gorm.DB) *handler.ArticleHandler {
	wire.Build(
		datastore.NewArticleDatastore,
		datastore.NewArticleCategoryDatastore,
		interactor.NewArticleInteractor,
		presenter.NewArticlePresenter,
		handler.NewArticleHandler,
	)
	return nil
}

func InitArticleTestHandler() *handler.ArticleHandler {
	wire.Build(
		interactor.NewArticleMockInteractor,
		presenter.NewArticleMockPresenter,
		handler.NewArticleHandler,
	)
	return nil
}
