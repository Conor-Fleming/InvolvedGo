//Clock package gives us functions to manipulate clocks and time
package clock

import "fmt"

type Clock struct {
	Hours   int
	Minutes int
}

func (clock Clock) String() string {
	return fmt.Sprintf("%.2v:%.2v", clock.Hours, clock.Minutes)
}

//New creates a new clock with two input values for hour and minute
func New(hour int, minute int) Clock {
	//manipulting values to ensure correct hour and minute values
	hour = (hour + (minute / 60)) % 24
	minute = minute % 60

	if minute < 0 {
		hour--
		minute = 60 + minute
	}
	if hour < 0 {
		hour = 24 + hour
	}

	myClock := Clock{
		Hours:   hour,
		Minutes: minute,
	}
	return myClock
}

//Add lets you add values to a clock
func (clock Clock) Add(minutes int) Clock {
	return New(clock.Hours, clock.Minutes+minutes)

}

//Subtract lets you take minutes from an existing clock
func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.Hours, clock.Minutes-minutes)
}
