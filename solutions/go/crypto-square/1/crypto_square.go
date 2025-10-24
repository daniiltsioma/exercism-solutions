package cryptosquare

import "unicode"


func Encode(pt string) string {
	input := []rune{}

	// normalize input: 
	// leave lowercase characters and digits only.
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			input = append(input, unicode.ToLower(r))
		}
	}

	// find dimensions for the rectangle.
	// start by setting both to 1.			
	c := 1
	r := 1
	// increment both until we get enough area...
	for c * r < len(input) {
		c, r = c+1, r+1
	}
	// try decreasing r if possible...
	if c * (r-1) >= len(input) {
		r -= 1
	}
	// now we got the right dimensions.

	// create a temporary matrix of characters
	// to represent r "lines" of the coded message.
	coded := make([][]rune, c)
	for i := range c {
		coded[i] = make([]rune, r)
	}
	
	// populate the newly created matrix.
	cr := c * r
	l := len(input)
	for i := 0; i < cr; i++ {
		j := i % c 		// current row
		k := i / c		// current column
		if i < l {
			coded[j][k] = input[i]
		} else {
			coded[j][k] = ' '
		}
	}

	// concatenate matrix into a single coded message.
	msg := []rune{}
	for i, row := range coded {
		msg = append(msg, row...)
		if i < len(coded) - 1 {
			msg = append(msg, ' ')
		} 
	}

	return string(msg)
}
