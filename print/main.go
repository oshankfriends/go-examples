package main

import (
	"fmt"
	"strings"
)

const (
	WeatherText        = "%s"
	City               = "%s"
	CurrentWeatherText = "%s"
	CurrentTemp        = "%d"
	LowTemp            = "%d"
	HighTemp           = "%d"
	Day                = "%s"
)

func main() {
	fmt.Println("Documentation is for users.")
	var a = [...]int{7, 6, 7, 6}
	fmt.Println(a)
	fmt.Printf("Today's temperature in "+City+" is "+WeatherText+" high of "+HighTemp+" degrees and low of "+LowTemp+" degrees.\n", "bengaluru", "cloudy", 23, 20)
	count := strings.Count("Today's temperature in "+City+" is "+WeatherText+" high of "+HighTemp+" degrees and low of "+LowTemp+" degrees.","%")
	fmt.Println(count)
}
