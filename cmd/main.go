package main

import (
	"fmt"
	"gocron-sample/database"
	"gocron-sample/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	database.OpenDb("postgres://root:root@localhost:5432/data-persistence?sslmode=disable")

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(2).Seconds().Do(service.PrintJob)

	fmt.Println("Starting scheduler...")
	scheduler.StartAsync()
	fmt.Println("Scheduler started...")

	// Exit handling
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	fmt.Println("Exiting scheduler app...")
}
