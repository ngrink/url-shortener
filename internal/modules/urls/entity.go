package urls

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	UserId      uint   `json:"user_id" db:"user_id"`
	Key         string `json:"key" db:"key"`
	OriginalURL string `json:"original_url" db:"original_url"`
	Visits      uint   `json:"visits" db:"visits"`
}
