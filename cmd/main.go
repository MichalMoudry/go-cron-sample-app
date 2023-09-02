package main

import (
	"fmt"
	"gocron-sample/database"
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
	database.OpenDb()
	database.CreateTables()

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(2).Seconds().Do(printJob)

	fmt.Println("Starting scheduler...")
	scheduler.StartAsync()
	fmt.Println("Scheduler started...")

	// Exit handling
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	fmt.Println("Exiting schduler app...")
}
