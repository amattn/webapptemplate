package main

import (
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	assertEqual(t, 186603930, time.Unix(internalBuildTimestamp, 0), BuildDate())
	assertEqual(t, 186603931, internalBuildNumber, BuildNumber())
	assertEqual(t, 186603932, internalVersionString, Version())
	VersionInfo()
}

func TestNothing(t *testing.T) {
	// do nothing.

	// uncomment the following line to verify test harness is working.

	// t.Error(3022615210)
}
