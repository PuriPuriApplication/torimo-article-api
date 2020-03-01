package repository

import "torimo-article-api/src/domain/model"

type IUserRepository interface {
	FindById(id uint64) model.User
}
