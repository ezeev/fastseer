package logger

import (
	"fmt"
	"log"
	"runtime"
)

func Info(shop string, message string) {
	log.Printf("INFO - Shop: %s, message: %s", shop, message)
}

func Debug(shop string, message string) {
	log.Printf("DEBUG - Shop: %s, message: %s", shop, message)
}

func Error(shop string, message string) {
	log.Printf("ERROR - Shop: %s, message: %s", shop, message)
	//trace
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("\t TRACE: %s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

func Trace(shop string, message string) {
	log.Printf("TRACE - Shop: %s, message: %s", shop, message)
	//trace
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("\t%s,:%d %s\n", frame.File, frame.Line, frame.Function)
}

func Fatal(shop string, message string) {
	log.Fatalf("FATAL - Shop: %s, message: %s", shop, message)
	//trace
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("\tTRACE: %s,:%d %s\n", frame.File, frame.Line, frame.Function)
}