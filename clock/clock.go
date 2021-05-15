package clock

import "fmt"

type Clock struct {
	Hours   int
	Minutes int
}

func (clock Clock) String() string {
	return fmt.Sprintf("%.2v:%.2v", clock.Hours, clock.Minutes)
}

func New(hour int, minute int) Clock {
	//manipulting values to ensure correct hour and minute values
	hour = (hour % 24) + ((minute / 60) % 24)
	minute = minute % 60

	if hour < 0 {

	}

	myClock := Clock{
		Hours:   hour,
		Minutes: minute,
	}
	return myClock
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.Hours, clock.Minutes+minutes)

}

func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.Hours, clock.Minutes-minutes)
}
