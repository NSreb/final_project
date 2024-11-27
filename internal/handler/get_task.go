package handler

import (
	"encoding/json"
	"go_final_project/internal/helper"
	"go_final_project/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetTask(w http.ResponseWriter, req *http.Request) {

	idStr := req.URL.Query().Get("id")

	if idStr == "" {
		helper.SendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	taskId := models.TaskId{Id: id}

	if taskId.Id == 0 {
		helper.SendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	result, err := h.repo.GetTask(taskId)

	if err != nil {
		http.Error(w, "Ошибка выполнения запроса", http.StatusInternalServerError)
		return
	}

	if result == nil {
		helper.SendJSONError(w, "Задача не найдена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Кодируем результат в JSON и отправляем его клиенту
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Ошибка кодирования данных в JSON", http.StatusInternalServerError)
		return
	}

}
