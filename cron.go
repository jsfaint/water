package main

import (
	"github.com/robfig/cron"
)

var c *cron.Cron

func cronTable(name string) {
	tables := loadCronFromFile(name)

	c = cron.New()
	for _, v := range tables {
		c.AddFunc(v, drinkWater)
	}

	c.Start()
}
