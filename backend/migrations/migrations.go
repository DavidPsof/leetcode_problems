package migrations

import (
	post "github.com/DavidPsof/leetcode_problems/backend/domain/post/postgre"
	user "github.com/DavidPsof/leetcode_problems/backend/domain/user/postgre"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.UserDB{}, &post.PostDB{}); err != nil {
		return err
	}

	return nil
}
