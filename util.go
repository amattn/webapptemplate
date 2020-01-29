package main

import (
	"log"
	"runtime"
	"time"
)

//
// Usage example: copy paste the two lines to the beginning of any function
// 	log.Println(current_function(), "entering")
//	defer trace(current_function(), time.Now())
func trace(functionName string, startTime time.Time) {
	elapsed := time.Since(startTime)
	log.Printf("%s exiting ~ %s elapsed", functionName, elapsed)
}

func currentFunction() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}
