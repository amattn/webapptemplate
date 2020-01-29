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

	"github.com/amattn/deeperror"
)

var showH bool
var showHelp bool
var showVersion bool

func init() {
	flag.BoolVar(&showH, "h", false, "show help message and exit(0)")
	flag.BoolVar(&showHelp, "help", false, "show help message and exit(0)")
	flag.BoolVar(&showVersion, "version", false, "show version info and exit(0)")
}

func main() {
	log.Println(currentFunction(), "entering")
	defer trace(currentFunction(), time.Now())

	// var err error
	log.Printf("Starting %v", VersionInfo())
	log.Printf("os.Args: %v", os.Args)
	log.Printf("Go (runtime:%v) (GOMAXPROCS:%d) (NumCPUs:%d)\n", runtime.Version(), runtime.GOMAXPROCS(-1), runtime.NumCPU())

	// command line flags:
	flag.Parse()

	if showVersion {
		os.Exit(0)
	}

	if showH || showHelp {
		flag.Usage()
		os.Exit(0)
	}

	hostAndPort := ":8080"
	log.Println("Starting HTTP server at", "\nhttp://"+hostAndPort)
	err := http.ListenAndServe(hostAndPort, getHandler())

	deeperror.Fatal(4074108258, "ListenAndServe returned error", err)
}

func getHandler() *http.ServeMux {
	serveMux := http.NewServeMux()
	// setup our handler
	serveMux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			custom404Handler(w, r)
			return
		}

		_, _ = fmt.Fprintf(w, "root!")
	})

	return serveMux
}

func custom404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprintf(w, "404 Not Found, %v", html.EscapeString(r.URL.Path))
}
