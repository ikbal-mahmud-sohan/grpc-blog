package post

import (
	"myGo/Blogs/blog/storage"
	"myGo/Blogs/blog/storage/postgres"
	"context"
)

type CoreSvc struct {
	core *postgres.Storage
}

func NewCoreSvc(s *postgres.Storage) *CoreSvc {
	return &CoreSvc{
		core: s,
	}
}

func (cs CoreSvc) Create(ctx context.Context, t storage.Post) (int64, error) {

	return cs.core.Create(ctx, t)
}
func (cs CoreSvc) List(ctx context.Context) ([]storage.Post, error) {

	return cs.core.List(ctx)
}
func (cs CoreSvc) Get(ctx context.Context, id int64) (storage.Post, error) {

	return cs.core.Get(ctx, id)
}
func (cs CoreSvc) Update(ctx context.Context, t storage.Post) error {

	return cs.core.Update(ctx, t)
}
func (cs CoreSvc) Delete(ctx context.Context, id int64) error {

	return cs.core.Delete(ctx, id)
}
