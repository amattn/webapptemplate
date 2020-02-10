package main

import (
	"log"
	"net/http"
	"testing"
)

func TestRoot(t *testing.T) {
	log.Println(1629850031)
	res, err := http.Get(globalTestServer.URL)
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 3246726949, 200, res.StatusCode, globalTestServer.URL)
}

func Test404(t *testing.T) {
	log.Println(1629850031)
	res, err := http.Get(globalTestServer.URL + "/not_a_real_path")
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 1177791974, 404, res.StatusCode, globalTestServer.URL)
}

func Test5XX(t *testing.T) {
	log.Println(1629850031)
	res, err := http.Get(globalTestServer.URL + "/test/500")
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 2774240274, 500, res.StatusCode, globalTestServer.URL)
}
