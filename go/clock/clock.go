package clock

import "fmt"

// Clock type is an int which stores the number of minutes elapsed in the day
type Clock int

const numberOfHours = 24
const numberOfMinutes = 60

//New creates a clock with specified hour and minutes
func New(hour, minute int) Clock {
	time := (hour*numberOfMinutes + minute) % (numberOfMinutes * numberOfHours)
	if time < 0 {
		time += numberOfMinutes * numberOfHours
	}
	return Clock(time)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/numberOfMinutes, c%numberOfMinutes)
}

//Add adds time to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

//Subtract subtracts time from the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}
