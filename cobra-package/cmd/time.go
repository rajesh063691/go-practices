package cmd

import "time"

func getTimeInTimezone(timezone string) (string, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}

	t := time.Now().In(location).Format(time.RFC1123)
	return t, nil
}
