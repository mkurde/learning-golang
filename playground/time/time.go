package main

import (
	"fmt"
	"time"
)

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02T15:04:05Z"
	startDate  = "2020-09-07T08:00:00Z"
	endDate    = "2020-09-12T00:00:00Z"
)

func main() {
	getDays()
	firstDayLastMonth()
	daysLastMonth()
}

func getDays() {
	location, _ := time.LoadLocation("Europe/Berlin")
	startDateParsed, err := time.Parse(timeFormat, startDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	startDateParsed = time.Date(startDateParsed.Year(), startDateParsed.Month(), startDateParsed.Day(), 0, 0, 0, 0, location)

	endDateParsed, err := time.Parse(timeFormat, endDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	endDateParsed = time.Date(endDateParsed.Year(), endDateParsed.Month(), endDateParsed.Day(), 0, 0, 0, 0, location)

	hours := endDateParsed.Sub(startDateParsed).Hours() / 24
	fmt.Printf("getDays: %d\n", int(hours))
}

func firstDayLastMonth() {
	location, _ := time.LoadLocation("Europe/Berlin")
	t := time.Now().In(location)
	t = t.AddDate(0, -1, 0)
	date := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, location)
	fmt.Printf("firstDayLastMonth: %v\n", date)
}

func daysLastMonth() {
	location, _ := time.LoadLocation("Europe/Berlin")
	t := time.Now().In(location)
	t = t.AddDate(0, -1, 0)
	year, month, _ := t.Date()
	days := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, location).Day()
	fmt.Printf("daysLastMonth: %d\n", days)
}
