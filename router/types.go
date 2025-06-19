// Package router реализует простой HTTP-маршрутизатор с поддержкой параметров пути.
package router

import (
	"net/http"
	"server/types"
)

// contextKey используется для хранения значений в контексте HTTP-запроса.
type contextKey string

// Route описывает один маршрут с шаблоном пути, сегментами и обработчиками по HTTP-методам.
type Route struct {
	pattern  string                       // шаблон маршрута, например "/beer/{id}"
	segments []string                     // сегменты пути, разбитые по "/"
	routes   map[string]types.HandlerFunc // обработчики для HTTP-методов (GET, POST и др.)
}

// Router содержит список маршрутов и внутренний HTTP-мультиплексор.
type Router struct {
	routes []Route        // список зарегистрированных маршрутов
	mux    *http.ServeMux // стандартный HTTP-мультиплексор для обработки запросов
}
