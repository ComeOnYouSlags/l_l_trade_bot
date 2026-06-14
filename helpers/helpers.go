package helpers

import (
	"time"
)

// GetTimestamp returns current UTC timestamp string with milliseconds
func GetTimestamp() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05.000")
}

// GetTimestampDateTime returns UTC date and time strings
func GetTimestampDateTime() (string, string) {
	now := time.Now().UTC()
	return now.Format("2006-01-02"), now.Format("15:04:05.000")
}

// GetTimestampString concatenates date and time
func GetTimestampString(date_point string, time_point string) string {
	return date_point + " " + time_point
}

// GetTimestampInt returns current UTC Unix timestamp in milliseconds
func GetTimestampInt() int64 {
	return time.Now().UTC().UnixMilli()
}

// IntTimeToString converts milliseconds timestamp to UTC formatted string
func IntTimeToString(int_time int64) string {
	t := time.UnixMilli(int_time).UTC()
	return t.Format("2006-01-02 15:04:05.000")
}

// StringTimeToInt converts UTC timestamp string to milliseconds int64
func StringTimeToInt(timeStr string) (int64, error) {
	layout := "2006-01-02 15:04:05.000"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, err
	}
	return t.UTC().UnixMilli(), nil
}

func SyncWithBybit() {
	// TODO!!!
	// Call GET /v5/market/time
	// Compare serverTime with local UTC time
	// Adjust if drift exceeds acceptable threshold (e.g., 500ms)
}
