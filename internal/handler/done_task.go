package handler

import (
	"go_final_project/internal/helper"
	"go_final_project/internal/models"
	"go_final_project/internal/next_date"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) DoneTask(w http.ResponseWriter, req *http.Request) {
	idStr := req.URL.Query().Get("id")

	if idStr == "" {
		helper.SendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.SendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	taskId := models.TaskId{Id: id}

	if taskId.Id == 0 {
		helper.SendJSONError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	result, err := h.repo.GetTask(taskId)

	if result.Repeat != "" {
		nextDate, err := next_date.NextDate(time.Now(), result.Date, result.Repeat)
		if err != nil {
			helper.SendJSONError(w, "Ошибка при вычислении следующей даты: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = h.repo.UpdDateTask(id, nextDate)

		if err != nil {
			helper.SendJSONError(w, "Ошибка при обновлении даты: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = h.repo.DeleteTask(id)

		if err != nil {
			helper.SendJSONError(w, "Ошибка при удалении задача: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))

}
