package clock

import (
	"fmt"
	"math"
	"strconv"
)

type Clock struct {
	time int
}

func New(hours int, mins int) Clock {
	hour := int(math.Mod(float64(hours), 24))
	delta := mins + (hour * 60)

	time := math.Mod(float64(delta), 1440)

	if time < 0 {
		time = 1440 + time
	}

	return Clock{int(time)}
}

func (c Clock) Add(mins int) Clock {
	return New(0, c.time+mins)
}

func (c Clock) Subtract(mins int) Clock {
	return New(0, c.time-mins)
}

func (c Clock) String() string {
	hourPadding := false
	minutePadding := false

	hours := c.time / 60
	if hours < 10 {
		hourPadding = true
	}

	minutes := math.Mod(float64(c.time), 60)
	if minutes < 10 {
		minutePadding = true
	}

	hourString := strconv.Itoa(hours)
	minuteString := strconv.Itoa(int(minutes))

	if hourPadding {
		hourString = fmt.Sprintf("0%s", hourString)
	}

	if minutePadding {
		minuteString = fmt.Sprintf("0%s", minuteString)
	}

	return fmt.Sprintf("%s:%s", hourString, minuteString)
}