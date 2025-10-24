package twelve

import "fmt"

func Verse(i int) string {
	days := []string{
		"first", "second", "third", "fourth",
		"fifth", "sixth", "seventh", "eighth",
		"ninth", "tenth", "eleventh", "twelfth",
	}

	gifts := []string{
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}

	v := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1])

	for j := i; j > 0; j-- {
		if j > 1 {
			v += fmt.Sprintf(" %s,", gifts[j-1])
		} else if i > 1 {
			v += fmt.Sprintf(" and %s.", gifts[0])
		} else {
			v += fmt.Sprintf(" %s.", gifts[0])
		}
	}

	return v
}

func Song() string {
	s := Verse(1)
	
	for i := 2; i <= 12; i++ {
		s += fmt.Sprintf("\n%s", Verse(i))
	}

	return s 
}
