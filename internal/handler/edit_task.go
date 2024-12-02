package handler

import (
	"encoding/json"
	"go_final_project/internal/helper"
	"go_final_project/internal/models"
	"net/http"
	"time"
)

func (h *Handler) EditTask(w http.ResponseWriter, req *http.Request) {
	var task models.Tasks

	if err := json.NewDecoder(req.Body).Decode(&task); err != nil {
		helper.SendJSONError(w, "Ошибка десериализации JSON:"+err.Error(), http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		helper.SendJSONError(w, "Заголовок задачи не указан", http.StatusBadRequest)
		return
	}

	if task.Repeat != "" {
		if !helper.IsValidRepeatFormat(task.Repeat) || task.Repeat == "m" || task.Repeat == "w" {
			helper.SendJSONError(w, "Правило повторения указано в неправильном формате или недопустимо", http.StatusBadRequest)
			return
		}
	}

	if task.Date == "" {
		task.Date = time.Now().Format("20060102") // Устанавливаем сегодняшнюю дату
	} else if !helper.IsValidDateFormat(task.Date) {
		helper.SendJSONError(w, "Дата представлена в неправильном формате, ожидается YYYYMMDD", http.StatusBadRequest)
		return
	}

	err := h.repo.EditTask(task)

	if err != nil {
		helper.SendJSONError(w, "Ошибка обновления задачи:"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
