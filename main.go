package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var show_h bool
var show_help bool
var show_version bool

func init() {
	flag.BoolVar(&show_h, "h", false, "show help message and exit(0)")
	flag.BoolVar(&show_help, "help", false, "show help message and exit(0)")
	flag.BoolVar(&show_version, "version", false, "show version info and exit(0)")
}

func main() {
	start := time.Now()
	defer func(start time.Time) {
		elapsed := time.Since(start)
		log.Printf("Total elapsed time ~ %s", elapsed)
	}(start)
	// var err error
	log.Printf("Starting recmetrics (%v, v%v, build %v)", runtime.Version(), Version(), BuildNumber())
	log.Printf("GOMAXPROCS (default:%d) (set to:%d)\n", runtime.GOMAXPROCS(runtime.NumCPU()), runtime.GOMAXPROCS(runtime.NumCPU()))

	// command line flags:
	flag.Parse()

	if show_version {
		os.Exit(0)
	}

	if show_h || show_help {
		flag.Usage()
		os.Exit(0)
	}

	// do stuff

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			custom404Handler(w, r)
			return
		}

		fmt.Fprintf(w, "root")
	})

	host_and_port := ":8080"
	log.Println("Starting HTTP server at", host_and_port)
	log.Fatal(http.ListenAndServe(host_and_port, nil))
}

func custom404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 Not Found, %v", html.EscapeString(r.URL.Path))
}