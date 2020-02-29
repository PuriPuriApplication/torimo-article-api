//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/handler"
	"torimo-article-api/src/infrastructure/db/store"
	"torimo-article-api/src/interactor"
)

func Initialize(d *gorm.DB) *handler.ArticleHandler {
	wire.Build(
		store.NewArticleDatastore,
		interactor.NewArticleInteractor,
		handler.NewArticleHandler,
	)
	return nil
}
