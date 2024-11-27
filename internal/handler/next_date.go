package handler

import (
	"fmt"
	"go_final_project/internal/next_date"
	"net/http"
	"time"
)

func (h *Handler) NextDate(w http.ResponseWriter, req *http.Request) {

	nowStr := req.FormValue("now")
	date := req.FormValue("date")
	repeat := req.FormValue("repeat")

	if repeat == "" {
		http.Error(w, fmt.Sprintf("Правило повторения не может быть пустым"), http.StatusBadRequest)
		return
	}

	now, err := time.Parse("20060102", nowStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Неверный формат даты 'now': %s. Ожидается формат YYYYMMDD.", nowStr), http.StatusBadRequest)
		return
	}

	nextDate, err := next_date.NextDate(now, date, repeat)
	if err != nil {
		fmt.Println("Ошибка:", err)
		http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем ошибку клиенту
	}

	w.Write([]byte(nextDate))
}
