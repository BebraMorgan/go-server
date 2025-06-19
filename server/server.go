// Package server содержит функции инициализации и запуска HTTP-сервера.
package server

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

// Init инициализирует роутер, регистрирует маршруты и запускает HTTP-сервер на порту 8000.
// В случае ошибки регистрации маршрутов или запуска сервера происходит логирование и завершение работы.
func Init() {
	apiRouter, err := router.InitRouter(routes)
	if err != nil {
		log.Fatalf("Error registering handlers: %v", err)
	}

	fmt.Println("Server started at http://localhost:8000")
	if err := http.ListenAndServe(":8000", apiRouter); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
