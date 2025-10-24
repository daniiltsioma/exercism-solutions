package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

// helper mod function
func mod(a, b int) int {
	return (a % b + b) % b
}

func New(h, m int) Clock {
	minutes := mod((h * 60 + m), (24 * 60))
	return Clock{
		h: minutes / 60,
		m: mod(minutes, 60),
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}