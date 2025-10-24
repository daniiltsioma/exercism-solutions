package house

var rhymes = []string{
	"the house that Jack built.",
	"the malt\nthat lay in ",
	"the rat\nthat ate ",
	"the cat\nthat killed ",
	"the dog\nthat worried ",
	"the cow with the crumpled horn\nthat tossed ",
	"the maiden all forlorn\nthat milked ",
	"the man all tattered and torn\nthat kissed ",
	"the priest all shaven and shorn\nthat married ",
	"the rooster that crowed in the morn\nthat woke ",
	"the farmer sowing his corn\nthat kept ",
	"the horse and the hound and the horn\nthat belonged to ",
}

func Verse(v int) string {
	output := "This is "
	for i := v-1; i >= 0; i-- {
		output += rhymes[i]
	}
	return output
}

func Song() string {
	output := Verse(1)
	for i := 2; i <= 12; i++ {
		output += "\n\n" + Verse(i)
	}
	return output
}
