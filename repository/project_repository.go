package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/agungdh/go-crud-api/entity"
)

// ProjectRepository adalah kontrak Data Mapper.
type ProjectRepository interface {
	Create(ctx context.Context, p *entity.Project) error
	Update(ctx context.Context, p *entity.Project) error
	DeleteByUUID(ctx context.Context, id uuid.UUID) error
	FindByUUID(ctx context.Context, id uuid.UUID) (*entity.Project, error)
	List(ctx context.Context, limit, offset int, q string) ([]*entity.Project, int64, error)
}

// GormProjectRepository implementasi dengan GORM (tanpa Active Record).
type GormProjectRepository struct {
	db *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) *GormProjectRepository {
	return &GormProjectRepository{db: db}
}

// Create: set timestamps di sini (bukan di entity).
func (r *GormProjectRepository) Create(ctx context.Context, p *entity.Project) error {
	now := time.Now().UTC()
	p.CreatedAt = &now
	p.UpdatedAt = &now

	// Pastikan UUID ada (kalau entity dibuat tanpa helper).
	if p.UUID == uuid.Nil {
		p.UUID = uuid.New()
	}

	return r.db.WithContext(ctx).Create(p).Error
}

func (r *GormProjectRepository) Update(ctx context.Context, p *entity.Project) error {
	now := time.Now().UTC()
	p.UpdatedAt = &now

	// Update by uuid untuk konsistensi, tapi tetap menjaga primary key kalau sudah ada.
	tx := r.db.WithContext(ctx).Model(&entity.Project{}).
		Where("uuid = ?", p.UUID).
		Updates(map[string]interface{}{
			"name":         p.Name,
			"description":  p.Description,
			"release_date": p.ReleaseDate,
			"updated_at":   p.UpdatedAt,
		})
	return tx.Error
}

func (r *GormProjectRepository) DeleteByUUID(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("uuid = ?", id).
		Delete(&entity.Project{}).Error
}

func (r *GormProjectRepository) FindByUUID(ctx context.Context, id uuid.UUID) (*entity.Project, error) {
	var p entity.Project
	if err := r.db.WithContext(ctx).
		Where("uuid = ?", id).
		First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *GormProjectRepository) List(ctx context.Context, limit, offset int, q string) ([]*entity.Project, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	var (
		items []*entity.Project
		total int64
		tx    = r.db.WithContext(ctx).Model(&entity.Project{})
	)
	if q != "" {
		tx = tx.Where("name ILIKE ? OR description ILIKE ?", "%"+q+"%", "%"+q+"%")
	}

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := tx.Order("created_at DESC").Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
