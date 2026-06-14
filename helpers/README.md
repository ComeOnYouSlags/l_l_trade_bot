# HELPERS

This package contains helper functions

### GetTimestamp

```go
func GetTimestamp() string       
```
args:           -

r.type:         string Timestamp ("YYYY-MM-DD HH:MM:SS.mil")

Returns current full timestamp with milliseconds

### GetTimestampDateTime

```go
func GetTimestampDateTime() (string, string)
```
args:           -

r.type:         string Date (YYYY-MM-DD), string Time (HH:MM:SS.mil)

Returns timestamp as Date and Time strings

### GetTimestampString

```go
func GetTimestampString(date_point string, time_point string) string
```
args:           string Date (YYYY-MM-DD), string Time (HH:MM:SS.mil)

r.type:         string Timestamp

Returns concatinated Date and Time into full timestamp string

### GetTimestampInt

```go
func GetTimestampInt() int64
```
args:           -

r.type:         int64 Timestamp

Returns current Unix timestamp in milliseconds

### IntTimeToString

```go
func IntTimeToString(int_time int64) string
```
args:           int64 Timestamp

r.type:         string Timestamp "YYYY-MM-DD HH:MM:SS.mil"

Returns a formatted string of converted milliseconds timestamp

### StringTimeToInt

```go
func StringTimeToInt(timeStr string) (int64, error)
```
args:           string Timestamp "YYYY-MM-DD HH:MM:SS.mil"

r.type:         int64 Timestamp

Returns an int64 value of a converted timestamp string