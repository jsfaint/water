package main

import (
	"flag"
)

type opts struct {
	group string
	time  string
}

func parseFlags() (opt opts) {
	flag.StringVar(&opt.time, "t", "time.txt", "cron table config")
	flag.StringVar(&opt.group, "g", "group.txt", "config for group names")
	flag.Parse()

	return
}
