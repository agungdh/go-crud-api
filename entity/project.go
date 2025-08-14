package entity

import (
	"time"

	"github.com/google/uuid"
)

// Project adalah domain entity murni (tanpa method persistence).
type Project struct {
	ID          int64      `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UUID        uuid.UUID  `gorm:"type:uuid;not null;column:uuid"     json:"uuid"`
	Name        string     `gorm:"type:varchar(255);not null;column:name" json:"name"`
	Description string     `gorm:"type:text;not null;column:description"  json:"description"`
	ReleaseDate *time.Time `gorm:"type:date;column:release_date"      json:"release_date,omitempty"`
	CreatedAt   *time.Time `gorm:"type:timestamptz;column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"type:timestamptz;column:updated_at" json:"updated_at,omitempty"`
}

// TableName override untuk GORM (tanpa menjadikan ini Active Record).
func (Project) TableName() string { return "project" }

// NewProject helper buat inisialisasi entity baru di layer atas (service / handler).
func NewProject(name, description string, releaseDate *time.Time) *Project {
	return &Project{
		UUID:        uuid.New(),
		Name:        name,
		Description: description,
		ReleaseDate: releaseDate,
	}
}
