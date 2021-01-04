package main

import (
	"fmt"
	"os"
	"time"
)

const (
	maxYears = 85 // fingers crossed.
	maxWeeks = maxYears * 52
	week     = 7 * day
	day      = 24 * time.Hour
)

func main() {
	now := time.Now()
	beginningOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	beginningOfNextMonth := beginningOfMonth.AddDate(0, 1, 0)

	birthdate, err := time.Parse("2006-01-02", os.Args[1])
	if err != nil {
		panic(err)
	}
	eol := birthdate.AddDate(maxYears, 0, 0)

	fmt.Printf("Birthdate taken as %+v.\n", birthdate)
	fmt.Printf("EOL expected       %+v at age %d.\n", eol, maxYears)

	for d := beginningOfMonth; beginningOfNextMonth.Sub(d) > 0; d = d.Add(week) {
		_, isoWeek := d.ISOWeek()
		sumWeeks := d.Sub(birthdate) / week
		percentage := float64(sumWeeks) / float64(maxWeeks) * 100.0
		fmt.Printf("Week #%d — %d/%d (%.2f%%)\n", isoWeek, sumWeeks, maxWeeks, percentage)
	}
}
