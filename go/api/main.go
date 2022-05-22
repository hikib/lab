package main

import (
	"fmt"

	"github.com/hikmet-kibar/lab/go/holidays/holiday"
	// "github.com/hikmet-kibar/lab/go-api/holiday"
)

func main() {
	countries := []string{"Romania", "Denmark", "Germany"}
	for _, country := range countries {
		fmt.Println("---")

		countryCode := holiday.GetCountryCode(country)
		fmt.Println(countryCode)

		nextHoliday := holiday.GetNextHoliday(countryCode)
		fmt.Println(nextHoliday.Date)
		fmt.Println(nextHoliday.Name)
		fmt.Println(nextHoliday.LocalName)
	}
}
