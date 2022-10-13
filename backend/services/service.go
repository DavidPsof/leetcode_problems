package services

import (
	"context"
	"github.com/DavidPsof/leetcode_problems/backend/pkg/api_models"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"mime/multipart"
)

type Service struct {
	User
	Post
}

type User interface {
	CreateUser(user api_models.User) error
}

type Post interface {
	CreatePost(ctx *context.Context, post *api_models.PostReq, file *multipart.FileHeader) error
	GetPost(ctx *context.Context, id int) (*api_models.PostResp, error)
	GetPosts() ([]api_models.PostResp, error)
}

func NewService(ctx context.Context, db *gorm.DB, minio *minio.Client) *Service {
	ps := NewPostService(db, minio)

	return &Service{
		Post: ps,
	}
}
