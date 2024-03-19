package timeutil

import (
	"time"
)

func ParseDayDateString(dateString string) (time.Time, error) {
	layout := "2006-01-02"

	parsedTime, err := time.Parse(layout, dateString)

	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
