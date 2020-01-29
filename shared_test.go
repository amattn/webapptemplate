package main

import (
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"
)

var (
	globalTestServer *httptest.Server
)

func init() {
	globalTestServer = httptest.NewServer(getHandler())
}

func assertEqual(t *testing.T, expected, candidate interface{}, printargs ...interface{}) {
	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == false {
		extra := fmt.Sprintln(printargs...)
		t.Errorf("Expected != Candidate\n%s\nExpected:\n%+v\nCandidate:\n%+v", extra, expected, candidate)
	}
}
