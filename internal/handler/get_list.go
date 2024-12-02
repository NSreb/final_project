package handler

import (
	"encoding/json"
	"go_final_project/internal/helper"
	"go_final_project/internal/repository"
	"net/http"
)

type Response struct {
	Tasks []repository.Data `json:"tasks"`
}

func (h *Handler) GetList(w http.ResponseWriter, req *http.Request) {

	results, err := h.repo.GetList(30)

	if err != nil {
		helper.SendJSONError(w, "Ошибка выполнения запрос"+err.Error(), http.StatusBadRequest)
		return
	}

	response := Response{Tasks: results}

	if len(results) == 0 {
		response.Tasks = []repository.Data{}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Кодируем результат в JSON и отправляем его клиенту
	if err := json.NewEncoder(w).Encode(response); err != nil {
		helper.SendJSONError(w, "Ошибка кодирования данных в JSON", http.StatusBadRequest)
		return
	}

}
