package services

import "github.com/jmoiron/sqlx"

type Service struct {
}

func NewService(db *sqlx.DB) *Service {
	return &Service{}
}
