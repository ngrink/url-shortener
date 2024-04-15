package urls

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	UserId      uint   `json:"user_id" db:"user_id"`
	Key         string `json:"key" db:"key"`
	ShortURL    string `json:"short_url" db:"short_url"`
	OriginalURL string `json:"original_url" db:"original_url"`
	Visits      uint   `json:"visits" db:"visits"`
}

type Visit struct {
	ID        uint      `json:"id" db:"id" gorm:"primarykey"`
	UrlId     uint      `json:"url_id" db:"url_id"`
	IpAddress string    `json:"ip_address" db:"ip_address"`
	UserAgent string    `json:"user_agent" db:"user_agent"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type VisitResponse struct {
	IpAddress string    `json:"ip_address" db:"ip_address"`
	UserAgent string    `json:"user_agent" db:"user_agent"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
