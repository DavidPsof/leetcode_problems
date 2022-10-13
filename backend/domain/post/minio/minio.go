package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type PostRepoMinIO struct {
	mc *minio.Client
}

func NewPostRepoMinIO(mc *minio.Client) *PostRepoMinIO {
	return &PostRepoMinIO{
		mc,
	}
}

func (r *PostRepoMinIO) Get() {

}

func (r *PostRepoMinIO) UploadFile(ctx context.Context, file *multipart.File, size int64, filename string) error {
	_, err := r.mc.PutObject(ctx, "posts", filename, *file, size, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
