package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func startScheduler() {
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create scheduler instance!")
		return
	}

	_, err = s.NewJob(
		gocron.DurationJob(time.Minute),
		gocron.NewTask(
			func() {
				db.DeleteExpiredMessages(time.Now())
			},
		),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to register scheduler task!")
		return
	}

	s.Start()
}
