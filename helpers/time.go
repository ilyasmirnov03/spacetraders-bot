package helpers

import "time"

func TimeDiffInSeconds(isoDateString string) (int64, error) {
	// Parse the input ISO date string
	parsedTime, err := time.Parse(time.RFC3339, isoDateString)
	if err != nil {
		return 0, err
	}

	// Calculate the time difference in seconds
	currentTime := time.Now().UTC()
	diffInSeconds := currentTime.Unix() - parsedTime.Unix()

	return diffInSeconds, nil
}
