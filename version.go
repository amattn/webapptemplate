package main

import (
	"fmt"
	"log"
	"time"
)

const (
	internalIdentifier     = "webapptemplate"
	internalBuildTimestamp = 1580925140
	internalBuildNumber    = 0
	internalVersionString  = "0.0.0"
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

func LogVersionInfo() {
	log.Printf(VersionInfo())
}
