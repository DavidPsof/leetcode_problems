package post

import "github.com/DavidPsof/leetcode_problems/backend/domain/post/postgre"

type PostRepository interface {
	Get(int) (*postgre.PostDB, error)
	GetAll() ([]postgre.PostDB, error)
	Create(name *postgre.PostDB) error
}
