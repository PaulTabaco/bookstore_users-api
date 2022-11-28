package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

func GetNowUTC() time.Time {
	return time.Now().UTC()
}

func GetNowUTCString() string {
	return GetNowUTC().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNowUTC().Format(apiDBLayout)
}
