package urls

import (
	"errors"
	"math/rand"
	"os"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrKeyExists = errors.New("key already exists")
)

type IUrlsService interface {
	CreateUrl(userId uint, data CreateUrlDto) (Url, error)
	GetAllUrls() ([]Url, error)
	GetUserUrls(userId uint64) ([]Url, error)
	GetUrl(id uint64) (Url, error)
	GetUrlByKey(key string) (Url, error)
	DeleteUrl(id uint64) error
	RegisterVisit(urlId uint, userAgent, ipAddress string) (Visit, error)
	GetUrlVisits(urlId uint64) ([]Visit, error)
	CheckOwner(userId, urlId uint64) (bool, error)
}

type UrlsService struct {
	repo IUrlsRepository
}

func NewUrlsService(repo IUrlsRepository) *UrlsService {
	return &UrlsService{
		repo: repo,
	}
}

func (s *UrlsService) CreateUrl(userId uint, data CreateUrlDto) (Url, error) {
	key := data.CustomKey

	if key != "" {
		if s.KeyExists(key) {
			return Url{}, ErrKeyExists
		}
	} else {
		for {
			key = generateShortKey()
			if !s.KeyExists(key) {
				break
			}
		}
	}

	url := Url{
		UserId:      userId,
		Key:         key,
		ShortURL:    os.Getenv("APP_HOST") + "/" + key,
		OriginalURL: data.OriginalUrl,
		Visits:      0,
	}

	var err error

	url, err = s.repo.CreateUrl(url)
	if err != nil {
		return Url{}, err
	}

	return url, nil

}

func (s *UrlsService) GetAllUrls() ([]Url, error) {
	urls, err := s.repo.GetAllUrls()
	if err != nil {
		return []Url{}, err
	}

	return urls, nil
}

func (s *UrlsService) GetUserUrls(userId uint64) ([]Url, error) {
	userUrls, err := s.repo.GetUserUrls(userId)
	if err != nil {
		return []Url{}, err
	}

	return userUrls, nil
}

func (s *UrlsService) GetUrl(id uint64) (Url, error) {
	url, err := s.repo.GetUrl(id)
	if err != nil {
		return Url{}, err
	}

	return url, nil
}

func (s *UrlsService) GetUrlByKey(key string) (Url, error) {
	url, err := s.repo.GetUrlByKey(key)
	if err != nil {
		return Url{}, err
	}

	return url, nil
}

func (s *UrlsService) DeleteUrl(id uint64) error {
	err := s.repo.DeleteUrl(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UrlsService) RegisterVisit(urlId uint, userAgent, ipAddress string) (Visit, error) {
	visit := Visit{
		UrlId:     urlId,
		UserAgent: userAgent,
		IpAddress: ipAddress,
	}

	visit, err := s.repo.RegisterVisit(visit)
	if err != nil {
		return Visit{}, err
	}

	return visit, nil
}

func (s *UrlsService) GetUrlVisits(urlId uint64) ([]Visit, error) {
	visits, err := s.repo.GetUrlVisits(urlId)
	if err != nil {
		return []Visit{}, err
	}

	return visits, nil
}

func (s *UrlsService) KeyExists(key string) bool {
	_, err := s.GetUrlByKey(key)
	return err == nil
}

func (s *UrlsService) CheckOwner(userId, urlId uint64) (bool, error) {
	url, err := s.repo.GetUrl(urlId)
	if err != nil {
		return false, ErrNotFound
	}

	if url.UserId != uint(userId) {
		return false, nil
	}

	return true, nil
}

func generateShortKey() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortLink := make([]byte, 6)

	for i := range shortLink {
		shortLink[i] = letters[rand.Intn(len(letters))]
	}

	return string(shortLink)
}
