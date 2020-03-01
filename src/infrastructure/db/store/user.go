package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"torimo-article-api/src/domain/model"
	"torimo-article-api/src/domain/repository"
)

type UserDatastore struct {
	db *gorm.DB
}

func NewUserDatastore(d *gorm.DB) repository.IUserRepository {
	return &UserDatastore{
		db: d,
	}
}

func (s *UserDatastore) FindById(id uint64) model.User {
	user := model.User{}
	s.db.Find(&user, "id=?", id)
	return user
}
