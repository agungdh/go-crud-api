package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/agungdh/go-crud-api/entity"
	"github.com/agungdh/go-crud-api/repository"
)

type ProjectService struct {
	db   *gorm.DB
	repo repository.ProjectRepository
}

func NewProjectService(db *gorm.DB, repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{db: db, repo: repo}
}

func (s *ProjectService) CreateProject(ctx context.Context, p *entity.Project) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := repository.NewGormProjectRepository(tx)
		return txRepo.Create(ctx, p)
	})
}
