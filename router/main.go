// Package router реализует простой HTTP-маршрутизатор с поддержкой параметров пути и методами регистрации обработчиков.
package router

import (
	"context"
	"net/http"
	"server/types"
)

const paramsKey contextKey = "params"

// InitRouter инициализирует новый Router, регистрирует маршруты через переданную функцию routes,
// и возвращает готовый к использованию маршрутизатор.
func InitRouter(routes func(r *Router)) (*Router, error) {
	newRouter := &Router{
		mux:    http.NewServeMux(),
		routes: make([]Route, 0),
	}
	routes(newRouter)
	for _, route := range newRouter.routes {
		rt := route

		newRouter.mux.HandleFunc(rt.pattern, func(w http.ResponseWriter, req *http.Request) {
			params, ok := matchAndExtractParams(rt.segments, req.URL.Path)
			if !ok {
				http.NotFound(w, req)
				return
			}

			handler, ok := rt.routes[req.Method]
			if !ok {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}

			ctx := context.WithValue(req.Context(), paramsKey, params)
			handler(w, req.WithContext(ctx))
		})
	}

	return newRouter, nil
}

// Get регистрирует обработчик для HTTP-метода GET по указанному пути.
func (r *Router) Get(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "GET")
}

// Post регистрирует обработчик для HTTP-метода POST по указанному пути.
func (r *Router) Post(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "POST")
}

// Put регистрирует обработчик для HTTP-метода PUT по указанному пути.
func (r *Router) Put(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "PUT")
}

// Patch регистрирует обработчик для HTTP-метода PATCH по указанному пути.
func (r *Router) Patch(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "PATCH")
}

// Delete регистрирует обработчик для HTTP-метода DELETE по указанному пути.
func (r *Router) Delete(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "DELETE")
}

// Options регистрирует обработчик для HTTP-метода OPTIONS по указанному пути.
func (r *Router) Options(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "OPTIONS")
}

// Head регистрирует обработчик для HTTP-метода HEAD по указанному пути.
func (r *Router) Head(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "HEAD")
}

// Connect регистрирует обработчик для HTTP-метода CONNECT по указанному пути.
func (r *Router) Connect(path string, handler types.JsonHandlerFunc) {
	RegisterRoute(r, path, handler, "CONNECT")
}

// ServeHTTP реализует интерфейс http.Handler и передаёт обработку запросов внутреннему mux.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
