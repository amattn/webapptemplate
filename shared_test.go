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

func assertEqual(t *testing.T, expected, candidate interface{}, printargs ...interface{}) {
	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == false {
		extra := fmt.Sprintln(printargs...)
		t.Errorf("Expected != Candidate\n%s\nExpected:\n%+v\nCandidate:\n%+v", extra, expected, candidate)
	}
}
