package clock

import "fmt"

type Clock 
New(hour, minute int) Clock

func CreateClock(hour int, minute int) string {
	New()
	hour = hour % 24
	tempMin := minute
	minute = minute % 60

	//use tempMin (initial number of minutes to update hour)
	if (tempMin / 60) > 24 {
		hour += tempMin/60 - 24
	} else {
		hour += tempMin / 60
	}

	fmt.Println(hour)
	fmt.Println(minute)

}

func AddMinutes() {

}

func SubtractMinutes() {

}

func SubtractMinutesStringless() {

}

func CompareClocks() {

}

func AddAndCompare() {

}

func SubtractAndCompare() {

}
