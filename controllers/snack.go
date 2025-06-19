// Package controllers содержит HTTP-обработчики для работы с сущностью Snack.
package controllers

import (
	"server/database"
	"server/models"
	"server/request"
	"server/types"
	"strconv"

	"gorm.io/gorm"
)

// CreateSnack создаёт новую запись закуски на основе JSON из запроса.
// Возвращает ошибку, если входные данные некорректны или произошла ошибка базы данных.
func CreateSnack(r *request.Request, params map[string]string) types.JsonResponse {
	var snack models.Snack
	if err := r.Json(&snack); err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid input: " + err.Error()}
	}

	if err := models.CreateSnack(database.DB, &snack); err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: snack}
}

// GetSnack возвращает закуску по ID, переданному в параметрах маршрута.
// Возвращает ошибку, если ID отсутствует, некорректен или запись не найдена.
func GetSnack(r *request.Request, params map[string]string) types.JsonResponse {
	idStr, ok := params["id"]
	if !ok {
		return types.JsonResponse{Status: "error", Message: "ID parameter required"}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid ID"}
	}

	snack, err := models.GetSnackByID(database.DB, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return types.JsonResponse{Status: "error", Message: "Snack not found"}
		}
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: snack}
}

// GetAllSnacks возвращает список всех закусок из базы данных.
// В случае ошибки возвращает соответствующее сообщение.
func GetAllSnacks(r *request.Request, params map[string]string) types.JsonResponse {
	snacks, err := models.GetAllSnacks(database.DB)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: snacks}
}

// UpdateSnack обновляет существующую запись закуски по ID.
// Возвращает ошибку, если ID отсутствует, некорректен, запись не найдена или входные данные некорректны.
func UpdateSnack(r *request.Request, params map[string]string) types.JsonResponse {
	idStr, ok := params["id"]
	if !ok {
		return types.JsonResponse{Status: "error", Message: "ID parameter required"}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid ID"}
	}

	snack, err := models.GetSnackByID(database.DB, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return types.JsonResponse{Status: "error", Message: "Snack not found"}
		}
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	var input models.Snack
	if err := r.Json(&input); err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid input: " + err.Error()}
	}

	snack.Name = input.Name
	snack.Type = input.Type
	snack.Description = input.Description
	snack.Country = input.Country
	snack.Calories = input.Calories
	snack.Spicy = input.Spicy
	snack.Vegetarian = input.Vegetarian

	if err := models.UpdateSnack(database.DB, snack); err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Data: snack}
}

// DeleteSnack удаляет запись закуски по ID.
// Возвращает ошибку, если ID отсутствует, некорректен или произошла ошибка при удалении.
func DeleteSnack(r *request.Request, params map[string]string) types.JsonResponse {
	idStr, ok := params["id"]
	if !ok {
		return types.JsonResponse{Status: "error", Message: "ID parameter required"}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return types.JsonResponse{Status: "error", Message: "Invalid ID"}
	}

	if err := models.DeleteSnack(database.DB, uint(id)); err != nil {
		return types.JsonResponse{Status: "error", Message: err.Error()}
	}

	return types.JsonResponse{Status: "success", Message: "Snack deleted"}
}

// GetRandomSnack возвращает случайную закуску из базы данных.
// Если закусок нет, возвращает соответствующее сообщение об ошибке.
func GetRandomSnack(r *request.Request, params map[string]string) types.JsonResponse {
	var snack models.Snack
	err := database.DB.Order("RAND()").First(&snack).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return types.JsonResponse{
				Status:  "error",
				Message: "No snacks found",
			}
		}
		return types.JsonResponse{
			Status:  "error",
			Message: err.Error(),
		}
	}

	return types.JsonResponse{
		Status: "success",
		Data:   snack,
	}
}
