package category

import (
	"myGo/Blogs/blog/storage"
	bgvc "myGo/Blogs/gunk/v1/category"
	"context"
)

type CategoryCoreLink interface{
	CreateCat(context.Context,storage.Category)(int64,error)
	ListCat(context.Context)([]storage.Category, error)
	GetCat(context.Context, int64)(storage.Category, error)
	UpdateCat(context.Context, storage.Category) error
	DeleteCat(context.Context,int64)error
}

type CategorySvc struct {
	bgvc.UnimplementedCategoryServiceServer
	store CategoryCoreLink
}

func NewCategorySvc(c CategoryCoreLink) *CategorySvc{
	return &CategorySvc{
		store: c,
	}
}
