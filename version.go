package main

import (
	"fmt"
	"time"
)

const (
	internalIdentifier           = "rootstrap"
	internalBuildTimestamp int64 = 1581275458
	internalBuildNumber    int64 = 10
	internalVersionString        = "0.0.1"
)

func BuildDate() time.Time {
	return time.Unix(internalBuildTimestamp, 0)
}
func BuildNumber() int64 {
	return internalBuildNumber
}
func Version() string {
	return internalVersionString
}

func VersionInfo() string {
	return fmt.Sprintf("%s (%v, build %v, build date:%v)", internalIdentifier, Version(), BuildNumber(), BuildDate())
}
