package article_repository

import "torimo-article-api/src/domain/article_model"

type IUserRepository interface {
	FindById(id uint64) article_model.User
}
