ackage main

import (
	"fmt"
	"time"
)

var version = "1.0.0"
var countryTz = map[string]string{
	"Eastern":  "America/New_York",
	"Central":  "America/Chicago",
	"London":   "Europe/London",
	"Sydney":   "Australia/Sydney",
	"Calcutta": "Asia/Calcutta",
}

func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func main() {
	utc := time.Now().UTC().Format("15:04")
	et := timeIn("Eastern").Format("15:04")
	ct := timeIn("Chicago").Format("15:04")
	eg := timeIn("London").Format("15:04")
	au := timeIn("Sydney").Format("15:04")
	it := timeIn("Calcutta").Format("15:04")

	fmt.Println("UTC Time: ", utc)
	fmt.Println("Eastern Time: ", et)
	fmt.Println("Central Time: ", ct)
	fmt.Println("England Time: ", eg)
	fmt.Println("Sydney Time : ", au)
	fmt.Println("Calcutta Time : ", it)

}
