package shorturl

import "gorm.io/gorm"

type IShortURLRepository interface {
	CreateShortURL(shorCode, url string) error
	GetShortURL(shortCode string) (*ShortURL, error)
	UpdateShortURL(shortCode, url string) error
	DeleteShortURL(shortCode string) error
}

type shortURLRepository struct {
	db *gorm.DB
}

func NewShortURLRepository(db *gorm.DB) *shortURLRepository {
	return &shortURLRepository{db: db}
}

func (r *shortURLRepository) CreateShortURL(shortCode, url string) error {
	shortURL := ShortURL{ShortCode: shortCode, OriginalURL: url}
	r.db.AutoMigrate(&ShortURL{})
	
	result:=r.db.Create(&shortURL)

	return result.Error
}

func (r *shortURLRepository) GetShortURL(shortCode string) (*ShortURL, error) {
	var shortURL ShortURL
	result := r.db.Where("short_code = ?", shortCode).First(&shortURL)
	if result.Error != nil {
		return nil, result.Error
	}
	db:=r.db.Model(&shortURL).Update("visited", shortURL.Visited+1)
	if db.Error != nil {
		return nil, db.Error
	}

	return &shortURL, result.Error
}

func (r *shortURLRepository) UpdateShortURL(shortCode, url string) error {
	result:=r.db.Model(&ShortURL{}).Where("short_code = ?", shortCode).Update("original_url", url)
	return result.Error
}

func (r *shortURLRepository) DeleteShortURL(shortCode string) error {
	result:=r.db.Model(&ShortURL{}).Where("short_code = ?", shortCode).Delete(&ShortURL{})
	return result.Error
}
