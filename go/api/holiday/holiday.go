package holiday

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Country struct {
	CountryCode string
	Name        string
}

func getCountryData() (countries []Country) {
	page := "https://date.nager.at/api/v3/AvailableCountries"
	resp, err := http.Get(page)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		panic(err)
	}
	return
}

func GetCountryCode(name string) (code string) {
	// OPTIONAL: make map with Name: CountryCode
	countries := getCountryData()
	// countries := make(map[string]string)
	for _, country := range countries {
		if country.Name != name {
			continue
		}
		code = country.CountryCode
		break
		// countries[v.Name] = v.CountryCode
	}
	// code = countries["Denmark"]
	return
}

type PublicHoliday struct {
	Date        string
	LocalName   string
	Name        string
	CountryCode string
	Fixed       bool
	Global      bool
	Counties    []string
	LaunchYear  int
	Types       []string
}

func getHolidayData(code string) (holidays []PublicHoliday) {
	page := fmt.Sprintf("https://date.nager.at/api/v3/NextPublicHolidays/%s", code)
	resp, err := http.Get(page)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&holidays); err != nil {
		panic(err)
	}
	return
}

func GetNextHoliday(code string) (holiday PublicHoliday) {
	holidays := getHolidayData(code)
	holiday = holidays[0]
	// holidayName = holidays[0].Date
	return
}
