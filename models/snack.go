// Package models содержит определения моделей данных и функции для работы с ними.
package models

import (
	"gorm.io/gorm"
	"time"
)

// Snack представляет модель закуски с основными характеристиками.
type Snack struct {
	ID        uint           `gorm:"primaryKey"` // Уникальный идентификатор
	CreatedAt time.Time      // Время создания записи
	UpdatedAt time.Time      // Время последнего обновления записи
	DeletedAt gorm.DeletedAt `gorm:"index"` // Время удаления записи (soft delete)

	Name        string `gorm:"type:varchar(100);not null"` // Название закуски
	Type        string `gorm:"type:varchar(50)"`           // Тип закуски
	Description string `gorm:"type:text"`                  // Описание закуски
	Country     string `gorm:"type:varchar(50)"`           // Страна происхождения
	Calories    int    `gorm:""`                           // Калорийность
	Spicy       bool   `gorm:""`                           // Острая ли закуска
	Vegetarian  bool   `gorm:""`                           // Вегетарианская ли закуска
}

// CreateSnack сохраняет новую запись закуски в базе данных.
func CreateSnack(db *gorm.DB, snack *Snack) error {
	return db.Create(snack).Error
}

// GetSnackByID возвращает закуску по её идентификатору.
func GetSnackByID(db *gorm.DB, id uint) (*Snack, error) {
	var snack Snack
	if err := db.First(&snack, id).Error; err != nil {
		return nil, err
	}
	return &snack, nil
}

// GetAllSnacks возвращает список всех закусок из базы данных.
func GetAllSnacks(db *gorm.DB) ([]Snack, error) {
	var snacks []Snack
	if err := db.Find(&snacks).Error; err != nil {
		return nil, err
	}
	return snacks, nil
}

// UpdateSnack обновляет существующую запись закуски в базе данных.
func UpdateSnack(db *gorm.DB, snack *Snack) error {
	return db.Save(snack).Error
}

// DeleteSnack удаляет запись закуски по идентификатору.
func DeleteSnack(db *gorm.DB, id uint) error {
	return db.Delete(&Snack{}, id).Error
}
