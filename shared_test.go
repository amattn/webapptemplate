package main

import (
	"fmt"
	"log"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/amattn/deeperror"
)

const (
	TestSessionKey = "TESTCi6Eq0m5J4t4z6z356mTsh3VD4xcKRoOpIX4LzcQY792"
)

var (
	globalTestServer *httptest.Server
)

func init() {
	log.Println(2683175790)
	path := "/assets"
	am, err := NewAssetManager(path)
	if err != nil {
		derr := deeperror.New(4147351464, currentFunction()+" Failure: NewAssetManager", err)
		derr.AddDebugField("path", path)
		log.Fatal(derr)
	}

	globalTestServer = httptest.NewServer(GetRouter(am, TestSessionKey))
	log.Println(2683175791)
}

func assertEqual(t *testing.T, debugNum int64, expected, candidate interface{}, printArgs ...interface{}) {
	if expected == nil {
		if candidate != nil {
			extra := fmt.Sprintln(printArgs...)
			t.Errorf("%d Expected != Candidate, Candidate should be nil\n%s\nExpected (%T):\n%+v\nCandidate (%T):\n%+v", debugNum, extra, expected, expected, candidate, candidate)
			return
		}
	}

	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == false {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d Expected != Candidate\n%s\nExpected (%T):\n%+v\nCandidate (%T):\n%+v", debugNum, extra, expected, expected, candidate, candidate)
	}
}

func assertNotEqual(t *testing.T, debugNum int64, expected, candidate interface{}, printArgs ...interface{}) {
	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == true {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d Expected == Candidate but we want !=\n%s\nExpected (%T):\n%+v\nCandidate (%T):\n%+v", debugNum, extra, expected, expected, candidate, candidate)
	}
}
