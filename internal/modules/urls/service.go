package urls

import (
	"math/rand"
)

type IUrlsService interface {
	CreateUrl(userId uint, data CreateUrlDto) (Url, error)
	GetAllUrls() ([]Url, error)
	GetUserUrls(userId uint64) ([]Url, error)
	GetUrl(id uint64) (Url, error)
	GetUrlByKey(key string) (Url, error)
	DeleteUrl(id uint64) error
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
	key := generateShortKey()
	url := Url{
		UserId:      userId,
		Key:         key,
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

func generateShortKey() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortLink := make([]byte, 6)

	for i := range shortLink {
		shortLink[i] = letters[rand.Intn(len(letters))]
	}

	return string(shortLink)
}
