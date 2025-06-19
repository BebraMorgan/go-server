// Package main содержит точку входа в приложение сервера для работы с пивом и закусками.
// В функции main происходит инициализация базы данных, автоматическая миграция моделей и запуск сервера.
package main

import (
	"log"
	"server/database"
	"server/models"
	"server/router"
	"server/server"
)

var apiRouter router.Router

// main инициализирует соединение с базой данных, выполняет миграцию моделей Beer и Snack,
// а затем запускает HTTP-сервер.
func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	database.DB.AutoMigrate(&models.Beer{})
	database.DB.AutoMigrate(&models.Snack{})

	server.Init()
}
