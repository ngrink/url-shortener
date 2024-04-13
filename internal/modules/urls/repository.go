package urls

import (
	"fmt"

	"gorm.io/gorm"
)

type IUrlsRepository interface {
	CreateUrl(url Url) (Url, error)
	GetAllUrls() ([]Url, error)
	GetUserUrls(userID uint64) ([]Url, error)
	GetUrl(id uint64) (Url, error)
	GetUrlByKey(key string) (Url, error)
	DeleteUrl(id uint64) error
}

/*
----------------------------------------------------------------------------

	SQL Repository

----------------------------------------------------------------------------
*/
type UrlsSQLRepository struct {
	db *gorm.DB
}

func NewUrlsSqlRepository(db *gorm.DB) *UrlsSQLRepository {
	db.AutoMigrate(&Url{})

	return &UrlsSQLRepository{db: db}
}

func (r *UrlsSQLRepository) CreateUrl(url Url) (Url, error) {
	result := r.db.Create(&url)
	if result.Error != nil {
		return Url{}, result.Error
	}

	return url, nil
}

func (r *UrlsSQLRepository) GetAllUrls() ([]Url, error) {
	urls := []Url{}
	result := r.db.Find(&urls)
	if result.Error != nil {
		return []Url{}, result.Error
	}

	return urls, nil
}

func (r *UrlsSQLRepository) GetUserUrls(userId uint64) ([]Url, error) {
	urls := []Url{}
	result := r.db.Where("user_id = ?", userId).Find(&urls)
	if result.Error != nil {
		return []Url{}, result.Error
	}

	return urls, nil
}

func (r *UrlsSQLRepository) GetUrl(id uint64) (Url, error) {
	var url Url
	result := r.db.Find(&url, id)
	if result.RowsAffected == 0 {
		return Url{}, fmt.Errorf("Url not found")
	}

	return url, nil
}

func (r *UrlsSQLRepository) GetUrlByKey(key string) (Url, error) {
	var url Url
	result := r.db.Where("key = ?", key).Find(&url)
	if result.RowsAffected == 0 {
		return Url{}, fmt.Errorf("Url not found")
	}

	return url, nil
}

func (r *UrlsSQLRepository) DeleteUrl(id uint64) error {
	url := Url{}
	result := r.db.Find(&url, id)
	if result.RowsAffected == 0 {
		return fmt.Errorf("Url not found")
	}

	r.db.Delete(&url)

	return nil
}
