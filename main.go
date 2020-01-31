package main

import (
	"flag"
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

const (
	DefaultSessionKey = "PLLKCi6Eq0m5J4t4z6z356mTsh3VD4xcKRoOpIX4LzcQY792"
)

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

	path := "/assets"
	am, err := NewAssetManager(path)
	if err != nil {
		derr := deeperror.New(4147351464, currentFunction()+" Failure: NewAssetManager", err)
		derr.AddDebugField("path", path)
		log.Fatal(derr)
	}

	hostAndPort := ":8080"
	log.Println("Starting HTTP server at", "\nhttp://"+hostAndPort)
	err = http.ListenAndServe(hostAndPort, GetRouter(am, DefaultSessionKey))

	deeperror.Fatal(4074108258, "ListenAndServe returned error", err)
}
