package usecase

import (
	"torimo-article-api/src/handler/request"
)

type IArticleUsecase interface {
	Create(requestArticle *request.RequestArticle) uint64
}
