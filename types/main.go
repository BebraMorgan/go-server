package types

import (
	"net/http"
	"server/request"
)

// HandlerFunc определяет тип функции-обработчика HTTP-запросов.
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// JsonResponse представляет структуру JSON-ответа API.
type JsonResponse struct {
	Status  string `json:"status"`            // Статус ответа, например "success" или "error"
	Message string `json:"message,omitempty"` // Сообщение об ошибке или дополнительная информация
	Data    any    `json:"data,omitempty"`    // Данные ответа (может быть любого типа)
}

// JsonHandlerFunc определяет тип функции-обработчика, которая принимает
// кастомный запрос и параметры маршрута, возвращая JSON-ответ.
type JsonHandlerFunc func(r *request.Request, params map[string]string) JsonResponse
