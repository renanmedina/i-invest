package main

import (
	"github.com/go-co-op/gocron/v2"
)

func main() {
	scheduler, err := gocron.NewScheduler()

	if err != nil {
		panic(err)
	}

	_, err = scheduler.NewJob(
		gocron.CronJob("0 * * * *", false),
		gocron.NewTask(
			func() {},
		),
	)

	if err != nil {
		panic(err)
	}

	scheduler.Start()
}
