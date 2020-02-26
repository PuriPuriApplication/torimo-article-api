package registry

import (
	"github.com/pkg/errors"
	"go.uber.org/dig"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/repository"
	"torimo-article-api/src/infrastructure/db/store"
	"torimo-article-api/src/usecase"
	"torimo-article-api/src/interactor"
)

func New(d *gorm.DB) *dig.Container {
    c := dig.New()
    if err := c.Provide(func (r repository.IArticleRepository) usecase.IArticleUsecase {
		return interactor.NewArticleInteractor(r)
	}); err != nil {
        panic(errors.WithStack(err))
    }
    if err := c.Provide(func () repository.IArticleRepository {
		return store.NewRefrigeratorDatastore(d)
	}); err != nil {
        panic(errors.WithStack(err))
    }
    return c
}