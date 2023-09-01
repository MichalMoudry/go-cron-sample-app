package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
)

// A simple job that prints a string.
func printJob() {
	fmt.Printf("%s: Executed print job.\n", time.Now())
}

func main() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(2).Seconds().Do(printJob)

	fmt.Println("Starting scheduler...")
	scheduler.StartAsync()
	fmt.Println("Scheduler started...")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}
