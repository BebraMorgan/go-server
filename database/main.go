// Package database отвечает за инициализацию и настройку подключения к базе данных.
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// DB — глобальный экземпляр подключения к базе данных.
var DB *gorm.DB

// Init инициализирует подключение к базе данных MySQL с заданными параметрами,
// настраивает пул соединений и логирование.
// Возвращает ошибку в случае неудачи подключения.
func Init() error {
	dsn := "root:123@tcp(localhost:3306)/servergo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Настройка пула соединений с базой данных
	sqlDB.SetMaxIdleConns(10)           // Максимальное количество неиспользуемых соединений
	sqlDB.SetMaxOpenConns(100)          // Максимальное количество открытых соединений
	sqlDB.SetConnMaxLifetime(time.Hour) // Максимальное время жизни соединения

	DB = db
	return nil
}
