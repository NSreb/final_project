package helper

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"
)

func SendJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("20060102", date)
	return err == nil
}

func IsValidRepeatFormat(repeat string) bool {
	validRepeat := regexp.MustCompile("^(d \\d+|y)$")
	return validRepeat.MatchString(repeat)
}
