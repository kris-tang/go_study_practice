package main

import "fmt"

type Alarm struct {
	Time string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time: time,
		Sound: "Kloxon",
	}
	return a
}

func main()  {
	fmt.Printf("%+v\n", NewAlarm("08:22"))
}