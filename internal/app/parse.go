package app

import (
	"fmt"
	"time"
)

// Parse string (yyyy-mm-dd hh:mm:ss) to time
func parseDateToTime(input string, fallback time.Time) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		fmt.Println(err)
		return fallback, err
	}
	return t, nil
}
