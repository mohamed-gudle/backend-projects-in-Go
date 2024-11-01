package shorturl

import "gorm.io/gorm"

type IShortURLRepository interface {
	CreateShortURL(shorCode, url string) error
	// GetShortURL(shortCode string) (*ShortURL, error)
	// UpdateShortURL(shortCode, url string) error
	// DeleteShortURL(shortCode string) error
}

type shortURLRepository struct {
	db *gorm.DB
}

func NewShortURLRepository(db *gorm.DB) *shortURLRepository {
	return &shortURLRepository{db: db}
}

func (r *shortURLRepository) CreateShortURL(shortCode, url string) error {
	shortURL := ShortURL{ShortCode: shortCode, OriginalURL: url}
	createStudentTableSQL := `CREATE TABLE "short_urls" (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"short_code" varchar(255) NOT NULL,
		"original_url" varchar(255) NOT NULL,
		"created_at" datetime NOT NULL,
		"updated_at" datetime NOT NULL
	  );` // SQL Statement for Create Table
	r.db.Exec(createStudentTableSQL);
	
	result:=r.db.Table("short_urls").Create(&shortURL)

	return result.Error
}
