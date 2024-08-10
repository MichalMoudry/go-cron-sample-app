package main

import (
	"fmt"
	"gocron-sample/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	id := uuid.New()
	scheduler.NewJob(
		gocron.DurationJob(200*time.Millisecond),
		//gocron.CronJob("*/1 * * * * *", true),
		gocron.NewTask(service.PrintJob, id),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)

	log.Println("Starting scheduler")
	scheduler.Start()

	// Exit handling
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	fmt.Println("Exiting scheduler app...")

	if err = scheduler.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
