package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func main() {
	deadline := flag.String("deadline", "",
		"The deadline for the countdown timer in RFC3339 format (e.g. 2020-12-25T15:00:00+01:00)")
	flag.Parse()

	if *deadline == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	v, err := time.Parse(time.RFC3339, *deadline)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for range time.Tick(1 * time.Second) {
		remainingTime := getRemainingTime(v)

		if remainingTime.t <= 0 {
			fmt.Println("Countdown reached")
			break
		}

		fmt.Printf("Days: %02d Hours: %02d Minutes: %02d Seconds: %02d\n",
			remainingTime.d, remainingTime.h, remainingTime.m, remainingTime.s)
	}

}

func getRemainingTime(t time.Time) countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())

	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
