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
		return repository.NewGormProjectRepository(tx).Create(ctx, p)
	})
}

func (s *ProjectService) UpdateProject(ctx context.Context, p *entity.Project) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return repository.NewGormProjectRepository(tx).Update(ctx, p)
	})
}

func (s *ProjectService) DeleteProjectByUUID(ctx context.Context, uuidStr string) error {
	return s.repo.DeleteByUUID(ctx, repository.MustParseUUID(uuidStr))
}

func (s *ProjectService) GetProjectByUUID(ctx context.Context, uuidStr string) (*entity.Project, error) {
	return s.repo.FindByUUID(ctx, repository.MustParseUUID(uuidStr))
}

func (s *ProjectService) ListProjects(ctx context.Context, limit, offset int, q string) ([]*entity.Project, int64, error) {
	return s.repo.List(ctx, limit, offset, q)
}
