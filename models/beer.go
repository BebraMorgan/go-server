// Package models содержит определения моделей данных и функции для работы с ними.
package models

import (
	"gorm.io/gorm"
	"time"
)

// Beer представляет модель пива с основными характеристиками.
type Beer struct {
	ID        uint           `gorm:"primaryKey"` // Уникальный идентификатор
	CreatedAt time.Time      // Время создания записи
	UpdatedAt time.Time      // Время последнего обновления записи
	DeletedAt gorm.DeletedAt `gorm:"index"` // Время удаления записи (soft delete)

	Name        string  `gorm:"type:varchar(100);not null"` // Название пива
	Brewery     string  `gorm:"type:varchar(100)"`          // Пивоварня
	Style       string  `gorm:"type:varchar(50)"`           // Стиль пива
	Alcohol     float32 `gorm:"type:float"`                 // Содержание алкоголя (%)
	Description string  `gorm:"type:text"`                  // Описание пива
	IBU         int     `gorm:""`                           // Горечь (International Bitterness Units)
	EBC         int     `gorm:""`                           // Цвет (European Brewery Convention)
}

// CreateBeer сохраняет новую запись пива в базе данных.
func CreateBeer(db *gorm.DB, beer *Beer) error {
	return db.Create(beer).Error
}

// GetBeerByID возвращает пиво по его идентификатору.
func GetBeerByID(db *gorm.DB, id uint) (*Beer, error) {
	var beer Beer
	if err := db.First(&beer, id).Error; err != nil {
		return nil, err
	}
	return &beer, nil
}

// GetAllBeers возвращает список всех пивных записей.
func GetAllBeers(db *gorm.DB) ([]Beer, error) {
	var beers []Beer
	if err := db.Find(&beers).Error; err != nil {
		return nil, err
	}
	return beers, nil
}

// UpdateBeer обновляет существующую запись пива в базе данных.
func UpdateBeer(db *gorm.DB, beer *Beer) error {
	return db.Save(beer).Error
}

// DeleteBeer удаляет запись пива по идентификатору.
func DeleteBeer(db *gorm.DB, id uint) error {
	return db.Delete(&Beer{}, id).Error
}
