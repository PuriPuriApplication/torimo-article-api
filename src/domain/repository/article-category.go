package repository

type IArticleCategoryRepository interface {
	SaveCategoriesForArticle(articleID uint64, categoryIDs []uint64)
}
