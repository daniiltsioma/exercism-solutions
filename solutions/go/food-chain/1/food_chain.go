package foodchain

var verses = [][]string{
	{
		"I know an old lady who swallowed a fly.",
		"I don't know why she swallowed the fly. Perhaps she'll die.",
	},
	{
		"I know an old lady who swallowed a spider.\nIt wriggled and jiggled and tickled inside her.",
		"She swallowed the spider to catch the fly.",
	},
	{
		"I know an old lady who swallowed a bird.\nHow absurd to swallow a bird!",
		"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	},
	{
		"I know an old lady who swallowed a cat.\nImagine that, to swallow a cat!",
		"She swallowed the cat to catch the bird.",
	},
	{
		"I know an old lady who swallowed a dog.\nWhat a hog, to swallow a dog!",
		"She swallowed the dog to catch the cat.",
	},
	{
		"I know an old lady who swallowed a goat.\nJust opened her throat and swallowed a goat!",
		"She swallowed the goat to catch the dog.",
	},
	{
		"I know an old lady who swallowed a cow.\nI don't know how she swallowed a cow!",
		"She swallowed the cow to catch the goat.",
	},
	{
		"I know an old lady who swallowed a horse.\nShe's dead, of course!",
	},
}

func Verse(v int) string {
	if v == 1 {
		return verses[0][0] + "\n" + verses[0][1]
	}
	if v == len(verses) {
		return verses[len(verses)-1][0]
	}

	output := verses[v-1][0]
	for i := v-1; i >= 0; i-- {
		output += "\n" +  verses[i][1]
	}

	return output
}

func Verses(start, end int) string {
	res := ""
	for i := start; i <= end; i++ {
		res += Verse(i)
		if i < end {
			res += "\n\n"
		}
	}
	return res
}

func Song() string {
	return Verses(1, 8)
}
