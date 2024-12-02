package handler

import (
	"encoding/json"
	"fmt"
	"go_final_project/internal/helper"
	"go_final_project/internal/models"
	"go_final_project/internal/next_date"
	"log"
	"net/http"
	"time"
)

func (h *Handler) AddTask(w http.ResponseWriter, req *http.Request) {
	var task models.Task

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

	today := time.Now().Format("20060102")
	if task.Date < today && (task.Repeat == "" || task.Repeat == " ") {
		task.Date = today // Устанавливаем сегодняшнюю дату
	} else if task.Date < today && task.Repeat != "" {
		nextDate, err := next_date.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			helper.SendJSONError(w, "Ошибка при вычислении следующей даты: "+err.Error(), http.StatusBadRequest)
			return
		}
		task.Date = nextDate
	}

	id, err := h.repo.AddTask(task)

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"id": fmt.Sprintf("%d", id)}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
