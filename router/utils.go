// Package router реализует маршрутизацию HTTP-запросов с поддержкой параметров пути и JSON-ответов.
package router

import (
	"encoding/json"
	"net/http"
	"server/request"
	"server/types"
	"strings"
)

// baseRoute регистрирует маршрут с указанным HTTP-методом и обработчиком в маршрутизаторе.
// Если маршрут с таким путем уже существует, добавляет или обновляет обработчик для метода.
func baseRoute(r *Router, path string, handler types.HandlerFunc, method string) {
	segments := splitPath(path)
	for i := range r.routes {
		if r.routes[i].pattern == path {
			if r.routes[i].routes == nil {
				r.routes[i].routes = make(map[string]types.HandlerFunc)
			}
			r.routes[i].routes[method] = handler
			return
		}
	}

	r.routes = append(r.routes, Route{
		pattern:  path,
		segments: segments,
		routes:   map[string]types.HandlerFunc{method: handler},
	})
}

// RegisterRoute регистрирует маршрут с указанным HTTP-методом и обработчиком типа JsonHandlerFunc.
// Паника происходит, если переданный обработчик не соответствует типу JsonHandlerFunc.
func RegisterRoute(r *Router, path string, handler any, method string) {
	h, ok := handler.(types.JsonHandlerFunc)
	if !ok {
		panic("handler должен быть типа JsonHandlerFunc для маршрутов с параметрами")
	}
	baseRoute(r, path, JsonHandlerWrapper(h), method)
}

// matchAndExtractParams проверяет соответствие сегментов пути шаблону маршрута,
// извлекает параметры из пути и возвращает их в виде словаря.
// Возвращает false, если путь не соответствует шаблону.
func matchAndExtractParams(patternSegments []string, path string) (map[string]string, bool) {
	reqSegments := splitPath(path)
	if len(reqSegments) != len(patternSegments) {
		return nil, false
	}

	params := make(map[string]string, len(patternSegments))
	for i, segment := range patternSegments {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			key := segment[1 : len(segment)-1]
			params[key] = reqSegments[i]
		} else if segment != reqSegments[i] {
			return nil, false
		}
	}

	return params, true
}

// splitPath разбивает путь URL на сегменты, удаляя ведущие и конечные слеши.
func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return []string{}
	}
	return strings.Split(path, "/")
}

// GetParams извлекает параметры маршрута из контекста HTTP-запроса.
// Если параметры отсутствуют, возвращает пустую карту.
func GetParams(r *http.Request) map[string]string {
	params, ok := r.Context().Value(paramsKey).(map[string]string)
	if !ok {
		return map[string]string{}
	}
	return params
}

// JsonHandlerWrapper оборачивает JsonHandlerFunc в стандартный http.HandlerFunc,
// обеспечивая парсинг параметров, инициализацию запроса и отправку JSON-ответа.
func JsonHandlerWrapper(handler types.JsonHandlerFunc) types.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := GetParams(r)
		req := request.InitRequest(r)
		response := handler(req, params)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
