package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/article_repository"
)

type UserDatastore struct {
	db *gorm.DB
}

func NewUserDatastore(d *gorm.DB) article_repository.IUserRepository {
	return &UserDatastore{
		db: d,
	}
}

func (s *UserDatastore) FindById(id uint64) article_model.User {
	user := article_model.User{}
	s.db.Find(&user, "id=?", id)
	return user
}
