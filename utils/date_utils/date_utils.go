package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:052"
)

func GetNowUTC() time.Time {
	return time.Now().UTC()
}

func GetNowUTCString() string {
	return GetNowUTC().Format(apiDateLayout)
}
