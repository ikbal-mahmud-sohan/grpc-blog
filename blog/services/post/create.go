package post

import (
	"myGo/Blogs/blog/storage"
	bgv "myGo/Blogs/gunk/v1/post"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PostSvc) CreatePost(ctx context.Context, req *bgv.CreatePostRequest) (*bgv.CreatePostResponse, error) {
	//Needs to validate post

	post := storage.Post{
		Title:       req.GetPost().Title,
		Description: req.GetPost().Description,
		CategoryId:  req.GetPost().CategoryId,
		Image:       req.Post.Image,
	}
	id, err := s.store.Create(context.Background(), post)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create post: %s", err)
	}

	return &bgv.CreatePostResponse{
		ID: id,
	}, nil

}

func (s *PostSvc) ListPost(ctx context.Context, req *bgv.ListPostRequest) (*bgv.ListPostResponse, error) {
	//Needs to validate post
	res, err := s.store.List(context.Background())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get post: %s", err)
	}
	var ctl []*bgv.Post

	for _, value := range res {
		ctl = append(ctl, &bgv.Post{
			ID:           value.ID,
			Title:        value.Title,
			Description:  value.Description,
			IsCompleted:  value.IsCompleted,
			CategoryName: value.CategoryName,
			Image:        value.Image,
		})
	}
	fmt.Printf("%#v", ctl)
	return &bgv.ListPostResponse{
		Post: ctl,
	}, nil
}
func (s *PostSvc) GetPost(ctx context.Context, req *bgv.GetPostRequest) (*bgv.GetPostResponse, error) {
	//Needs to validate post
	res, err := s.store.Get(context.Background(), req.GetID())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get post: %s", err)
	}

	return &bgv.GetPostResponse{
		Post: &bgv.Post{
			ID:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			IsCompleted: res.IsCompleted,
			Image: res.Image,
			CategoryId: res.CategoryId,
			CategoryName: res.CategoryName,
		},
	}, nil
}

func (s *PostSvc) UpdatePost(ctx context.Context, req *bgv.UpdatePostRequest) (*bgv.UpdatePostResponse, error) {
	//Needs to validate post
	post := storage.Post{
		ID:          req.Post.ID,
		Title:       req.Post.Title,
		Description: req.Post.Description,
		CategoryId: req.Post.CategoryId,
		Image: req.Post.Image,

	}
	err := s.store.Update(context.Background(), post)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get post: %s", err)
	}

	return &bgv.UpdatePostResponse{}, nil

}

func (s *PostSvc) DeletePost(ctx context.Context, req *bgv.DeletePostRequest) (*bgv.DeletePostResponse, error) {
	//Needs to validate post

	err := s.store.Delete(context.Background(), req.GetID())

	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to Delete category.")
	}
	return &bgv.DeletePostResponse{}, nil
}
