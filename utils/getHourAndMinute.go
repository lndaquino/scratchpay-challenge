package utils

import (
	"errors"
	"strconv"
	"strings"
)

// GetHourAndMinute receives a time in the format HH:MM and returns hour and minute as *int
func GetHourAndMinute(time string) (*int, *int, error) {
	parsedTime := strings.Split(time, ":")

	if len(parsedTime) != 2 {
		return nil, nil, errors.New("Time must be in the format HH:MM")
	}

	hour, err := strconv.Atoi(parsedTime[0])
	if err != nil {
		return nil, nil, err
	}
	if hour < 0 || hour > 24 {
		return nil, nil, errors.New("Hour must be between 0 and 24")
	}

	minute, err := strconv.Atoi(parsedTime[1])
	if err != nil {
		return nil, nil, err
	}
	if minute < 0 || minute > 59 {
		return nil, nil, errors.New("Minute must be between 0 and 59")
	}

	return &hour, &minute, nil
}
