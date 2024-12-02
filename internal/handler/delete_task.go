package handler

import (
	"go_final_project/internal/helper"
	"go_final_project/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, req *http.Request) {
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

	err = h.repo.DeleteTask(id)

	if err != nil {
		helper.SendJSONError(w, "Ошибка при удалении задача: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{}"))
}
