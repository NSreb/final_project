package next_date

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {

	startDate, err := time.Parse("20060102", date)
	if err != nil {
		return "", fmt.Errorf("Неверный формат даты: %v", err)
	}

	var nextDate time.Time
	nextDate = startDate

	if strings.HasPrefix(repeat, "d ") {
		daysStr := strings.TrimSpace(repeat[2:]) // Извлекаем число после "d "
		days, err := strconv.Atoi(daysStr)
		if err != nil || days < 1 || days > 400 {
			return "", errors.New("Недопустимое количество дней в правиле ")
		}

		nextDate = nextDate.AddDate(0, 0, days)

		for !nextDate.After(now) {
			nextDate = nextDate.AddDate(0, 0, days)
		}
	} else if repeat == "y" {

		nextDate = nextDate.AddDate(1, 0, 0)

		for !nextDate.After(now) {
			nextDate = nextDate.AddDate(1, 0, 0)
		}
	} else {
		return "", errors.New("Не допустимый формат повторения")
	}
	return nextDate.Format("20060102"), nil
}
