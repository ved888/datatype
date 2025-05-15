package question

import (
	"strconv"
	"strings"
)

func DayOfYear(date string) int {
	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	split := strings.Split(date, "-")
	year, _ := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		daysInMonth[1] = 29
	}

	totalDays := 0
	for i := 0; i < month-1; i++ {
		totalDays += daysInMonth[i]
	}

	return totalDays + day
}
