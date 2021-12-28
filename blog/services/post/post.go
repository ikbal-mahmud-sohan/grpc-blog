package post

import (
	"myGo/Blogs/blog/storage"
	bgv "myGo/Blogs/gunk/v1/post"
	"context"
)

type PostCoreLink interface{
	Create(context.Context,storage.Post)(int64,error)
	List(context.Context)([]storage.Post, error)
	Get(context.Context, int64)(storage.Post, error)
	Update(context.Context, storage.Post) error
	Delete(context.Context,int64)error
}

type PostSvc struct {
	bgv.UnimplementedPostServiceServer
	store PostCoreLink
}

func NewPostSvc(s PostCoreLink) *PostSvc{
	return &PostSvc{
		store: s,
	}
}