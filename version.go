package main

import "time"

const (
	internal_BUILD_TIMESTAMP = 1500000000
	internal_BUILD_NUMBER    = 0
	internal_VERSION_STRING  = "2.0.0"
)

func BuildDate() time.Time {
	return time.Unix(internal_BUILD_TIMESTAMP, 0)
}
func BuildNumber() int64 {
	return internal_BUILD_NUMBER
}
func Version() string {
	return internal_VERSION_STRING
}
