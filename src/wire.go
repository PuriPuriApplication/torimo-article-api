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

func Initialize(d *gorm.DB) *handler.ArticleHandler {
	wire.Build(
		datastore.NewArticleDatastore,
		datastore.NewArticleCategoryDatastore,
		interactor.NewArticleInteractor,
		presenter.NewUserDataConvert,
		presenter.NewShopDataConvert,
		presenter.NewCategoryDataConvert,
		interactor.NewResponseInteractor,
		handler.NewArticleHandler,
	)
	return nil
}
