// Package controllers содержит HTTP-обработчики для работы с сущностью Beer.
package controllers

import (
	"server/database"
	"server/models"
	"server/request"
	"server/types"
	"strconv"

	"gorm.io/gorm"
)

// GetRandomBeer возвращает случайное пиво из базы данных.
// Если пиво не найдено, возвращает ошибку с соответствующим сообщением.
func GetRandomBeer(r *request.Request, params map[string]string) types.JsonResponse {
	var beer models.Beer
	err := database.DB.Order("RAND()").First(&beer).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return types.JsonResponse{Status: "error", Message: "No beers found"}
		}
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: beer}
}

// StoreBeer создаёт новую запись пива на основе JSON из запроса.
// Возвращает ошибку, если входные данные некорректны или произошла ошибка базы данных.
func StoreBeer(r *request.Request, params map[string]string) types.JsonResponse {
	var beer models.Beer
	if err := r.Json(&beer); err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid input: " + err.Error()}
	}

	if err := database.DB.Create(&beer).Error; err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: beer}
}

// ShowBeer возвращает пиво по ID, переданному в параметрах маршрута.
// Возвращает ошибку, если ID отсутствует, некорректен или запись не найдена.
func ShowBeer(r *request.Request, params map[string]string) types.JsonResponse {
	idStr, ok := params["id"]
	if !ok {
		return types.JsonResponse{Status: "error", Message: "ID parameter is required"}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid ID"}
	}

	var beer models.Beer
	if err := database.DB.First(&beer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return types.JsonResponse{Status: "error", Message: "Beer not found"}
		}
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: beer}
}

// UpdateBeer обновляет существующую запись пива по ID.
// Возвращает ошибку, если ID отсутствует, некорректен, запись не найдена или входные данные некорректны.
func UpdateBeer(r *request.Request, params map[string]string) types.JsonResponse {
	idStr, ok := params["id"]
	if !ok {
		return types.JsonResponse{Status: "error", Message: "ID parameter is required"}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid ID"}
	}

	var beer models.Beer
	if err := database.DB.First(&beer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return types.JsonResponse{Status: "error", Message: "Beer not found"}
		}
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	var input models.Beer
	if err := r.Json(&input); err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid input: " + err.Error()}
	}

	beer.Name = input.Name
	beer.Brewery = input.Brewery
	beer.Style = input.Style
	beer.Alcohol = input.Alcohol
	beer.Description = input.Description
	beer.IBU = input.IBU
	beer.EBC = input.EBC

	if err := database.DB.Save(&beer).Error; err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: beer}
}

// DeleteBeer удаляет запись пива по ID.
// Возвращает ошибку, если ID отсутствует, некорректен или произошла ошибка при удалении.
func DeleteBeer(r *request.Request, params map[string]string) types.JsonResponse {
	idStr, ok := params["id"]
	if !ok {
		return types.JsonResponse{Status: "error", Message: "ID parameter is required"}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid ID"}
	}

	if err := database.DB.Delete(&models.Beer{}, id).Error; err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Message: "Beer deleted"}
}
