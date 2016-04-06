package main

import (
	"flag"
	"fmt"
	"log"
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

	fmt.Println("hello world")
}
