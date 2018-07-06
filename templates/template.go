package main

import (
	"time"
	"html/template"
	"os"
)

type SpeakOut struct {
	Available            bool
	TempUnit             string
	TimeAsked            time.Time
	HighTemp             int
	LowTemp              int
	CurrentTemp          int
	RainProb             int32
	HoursOfRain          float64
	City                 string
	LocationTimeZone     string
	WeatherText          string
	CurrentWeatherText   string
	Day                  string
}

var tmpl = `{{if eq .Day "today" -}}
It will {{.WeatherText}} on {{.Day}}
{{else if eq .Day "tomorrow" -}}
It is {{.Day}}
{{else -}}
on {{.Day}} it will {{.WeatherText}}
{{end}}`

func main() {
	var s = &SpeakOut{
		RainProb:34,
		HighTemp:23,
		WeatherText:"cloudy and thunderstorm",
		Day:"tomorrow",
	}

	t := template.Must(template.New("test").Parse(tmpl))
	t.Execute(os.Stdout,s)
	
}
