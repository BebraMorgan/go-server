// Package server содержит маршруты инициализации HTTP-сервера.
package server

import (
	"server/controllers"
	"server/router"
)

// routes регистрирует все маршруты HTTP-сервера и связывает их с соответствующими контроллерами.
func routes(r *router.Router) {
	// Маршруты для работы с пивом (Beer)
	r.Get("/beer/random/", controllers.GetRandomBeer)
	r.Post("/beer/", controllers.StoreBeer)
	r.Get("/beer/{id}", controllers.ShowBeer)
	r.Put("/beer/{id}", controllers.UpdateBeer)
	r.Patch("/beer/{id}", controllers.UpdateBeer)
	r.Delete("/beer/{id}", controllers.DeleteBeer)

	// Маршруты для работы с закусками (Snack)
	r.Post("/snack/", controllers.CreateSnack)
	r.Get("/snack/random/", controllers.GetRandomSnack)
	r.Get("/snack/{id}", controllers.GetSnack)
	r.Get("/snacks/", controllers.GetAllSnacks)
	r.Put("/snack/{id}", controllers.UpdateSnack)
	r.Patch("/snack/{id}", controllers.UpdateSnack)
	r.Delete("/snack/{id}", controllers.DeleteSnack)

	// Дополнительный маршрут
	r.Get("/hohol/", controllers.GetRandomBeer)
}
