package services

import (
	"context"
	miniorepo "github.com/DavidPsof/leetcode_problems/backend/domain/post/minio"
	"github.com/DavidPsof/leetcode_problems/backend/domain/post/postgre"
	"github.com/DavidPsof/leetcode_problems/backend/pkg/api_models"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"mime/multipart"
)

// PostService service for post
type PostService struct {
	pgRepo    *postgre.PostRepoPostgre
	minioRepo *miniorepo.PostRepoMinIO
}

// NewPostService return new post service
func NewPostService(db *gorm.DB, mc *minio.Client) *PostService {
	pg := postgre.NewPostRepoPostgre(db)
	m := miniorepo.NewPostRepoMinIO(mc)

	return &PostService{
		pgRepo:    pg,
		minioRepo: m,
	}
}

// GetPost return post record and md file
func (s *PostService) GetPost(ctx *context.Context, id int) (*api_models.PostResp, error) {
	get, err := s.pgRepo.Get(id)
	if err != nil {
		return nil, err
	}

	resp := get.ToDomain()
	apiResp := resp.ToApiResp()

	return &apiResp, nil
}

// GetPosts return post records
func (s *PostService) GetPosts() ([]api_models.PostResp, error) {
	// TODO: пока возвращает все записи, потом нужно будет добавить топики и выдачу по подпискам

	posts, err := s.pgRepo.GetAll()
	if err != nil {
		return nil, err
	}

	agPosts := make([]api_models.PostResp, 0, len(posts))
	for i := range posts {
		toDomain := posts[i].ToDomain()
		agPosts = append(agPosts, toDomain.ToApiResp())
	}

	return agPosts, nil
}

// CreatePost insert post record and md file in minio
func (s *PostService) CreatePost(ctx *context.Context, post *api_models.PostReq, fh *multipart.FileHeader) error {
	postDB := new(postgre.PostDB)
	postDB.MinioUID = uuid.New().String()
	postDB.Title = post.Title

	file, err := fh.Open()
	if err != nil {
		return err
	}

	if err = s.minioRepo.UploadFile(*ctx, &file, fh.Size, postDB.MinioUID); err != nil {
		return err
	}

	if err = s.pgRepo.Create(postDB); err != nil {
		return err
	}

	return nil
}
