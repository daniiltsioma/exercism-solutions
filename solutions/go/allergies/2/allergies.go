package allergies

var knownAllergens = map[string]uint{
	"eggs": 1,
	"peanuts": 2,
	"shellfish": 4,
	"strawberries": 8,
	"tomatoes": 16,
	"chocolate": 32,
	"pollen": 64,
	"cats": 128,
}

func Allergies(allergies uint) []string {
	res := []string{}

	for k, v := range knownAllergens {
		if allergies & v == v {
			res = append(res, k)
		}
	}

	return res
}

func AllergicTo(allergies uint, allergen string) bool {	
	v, ok := knownAllergens[allergen]; if !ok {
		return false
	}
	return allergies & v == v
}
