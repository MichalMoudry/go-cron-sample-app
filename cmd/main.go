package main

import (
	"fmt"
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

	var input string
	for input != "exit" {
		fmt.Scanln(&input)
	}
}
