package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AndrewJDGalga/keep_awake/keepawake"
)

func main() {
	awake := keepawake.New()
	if err := awake.Start(); err != nil {
		log.Fatalf("Failed to start: %v\n", err)
	}

	log.Println("Keepawake is active.\nPress CTRL+C to quit.")

	//capture interrupt and attempt graceful exit, especially on Windows
	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, os.Interrupt, syscall.SIGTERM)
	<-sChan

	log.Println("Attempting graceful quit...")
	if err := awake.Stop(); err != nil {
		log.Fatalf("Warning: Failed to release lock! Error: %v\n", err)
	}
	log.Println("Success!")
}
