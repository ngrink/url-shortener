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
	RegisterVisit(visit Visit) (Visit, error)
	GetUrlVisits(urlId uint64) ([]Visit, error)
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
	db.AutoMigrate(
		&Url{}, &Visit{},
	)
	db.Migrator().DropColumn(&Visit{}, "updated_at")
	db.Migrator().DropColumn(&Visit{}, "deleted_at")

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
	result := r.db.Where("user_id = ?", userId).Order("created_at desc").Find(&urls)
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

	url.Visits += 1
	r.db.Save(&url)

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

func (r *UrlsSQLRepository) RegisterVisit(visit Visit) (Visit, error) {
	result := r.db.Create(&visit)
	if result.Error != nil {
		return Visit{}, result.Error
	}

	return visit, nil
}

func (r *UrlsSQLRepository) GetUrlVisits(urlId uint64) ([]Visit, error) {
	visits := []Visit{}
	result := r.db.Where("url_id = ?", urlId).Find(&visits)
	if result.Error != nil {
		return []Visit{}, result.Error
	}

	return visits, nil
}
