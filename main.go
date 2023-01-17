package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// log to custom file
	LOG_FILE := "/tmp/event-listener.log"

	// open log file
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// set the log to multiple output
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	var count uint64 = 0
	for count < 10 {
		count += 1
		log.Println("Counting", count)
		time.Sleep(time.Second)
	}
	log.Println("Disconnecting...")
	for {
		count += 1
		log.Println("Still counting", count)
		time.Sleep(time.Second)
	}
}
