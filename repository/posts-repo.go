package repository

import (
	"course/entity"
)

type PostRepository interface {
	Save(p *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
