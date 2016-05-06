package main

import (
	"log"
	"net/http"
	"testing"
)

func TestRoot(t *testing.T) {
	res, err := http.Get(test_server.URL)
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 200, res.StatusCode, test_server.URL)
}
