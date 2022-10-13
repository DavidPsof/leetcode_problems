package postgre

import (
	"github.com/DavidPsof/leetcode_problems/backend/aggregate"
	"gorm.io/gorm"
)

type PostDB struct {
	Title    string
	MinioUID string
	gorm.Model
}

type PostRepoPostgre struct {
	db *gorm.DB
}

func NewPostRepoPostgre(db *gorm.DB) *PostRepoPostgre {
	return &PostRepoPostgre{
		db,
	}
}

// Get return one record
func (r *PostRepoPostgre) Get(id int) (*PostDB, error) {
	var res PostDB

	r.db.First(&res, id)

	return &res, nil
}

// GetAll return all records
func (r *PostRepoPostgre) GetAll() ([]PostDB, error) {
	res := make([]PostDB, 0)

	r.db.Find(&res)

	return res, nil
}

// Create add new post record to postgre
func (r *PostRepoPostgre) Create(post *PostDB) error {
	res := r.db.Create(&post)

	return res.Error
}

func (d *PostDB) ToDomain() aggregate.Post {
	return aggregate.Post{
		Title: d.Title,
		Link:  d.MinioUID,
	}
}
