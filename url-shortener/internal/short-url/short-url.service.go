package shorturl

type IShortURLService interface {
	CreateShortURL(shortCode, url string) error
	GetShortURL(shortCode string) (*ShortURL, error)
	UpdateShortURL(shortCode, url string) error
	DeleteShortURL(shortCode string) error
}

type defaultShortURLService struct {
	repo IShortURLRepository
}

func NewDefaultShortURLService(repo IShortURLRepository) *defaultShortURLService {
	return &defaultShortURLService{repo: repo}
}

func (s *defaultShortURLService) CreateShortURL(shortCode, url string) error {
	return s.repo.CreateShortURL(shortCode, url)
}

func (s *defaultShortURLService) GetShortURL(shortCode string) (*ShortURL, error) {
	return s.repo.GetShortURL(shortCode)
}

func (s *defaultShortURLService) UpdateShortURL(shortCode, url string) error {
	return s.repo.UpdateShortURL(shortCode, url)
}

func (s *defaultShortURLService) DeleteShortURL(shortCode string) error {
	return s.repo.DeleteShortURL(shortCode)
}
