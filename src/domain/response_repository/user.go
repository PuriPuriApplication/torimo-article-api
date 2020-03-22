package response_repository

import (
	"torimo-article-api/src/domain/article_model"
	"torimo-article-api/src/domain/response_model"
)

type IUserRepository interface {
	Convert(user *article_model.User) *response_model.ResponseUser
}
