package main

import (
	"github.com/fatih/structs"
	"fmt"
	"strings"
)

type Event struct {
	Name     string
	Location *Location
	Level    string
}

type Location struct{
	Longitude float64
	Latitude  float64
}

func main() {
	e := &Event{
		Name:"test_event",
		Location:&Location{
			Longitude:3.87,
			Latitude:4.65,
		},
		Level:"event",
	}
	m := structs.Map(e)
	fmt.Printf("%v\n",m)
	fmt.Println(strings.ToLower("Hello_World"))
}
