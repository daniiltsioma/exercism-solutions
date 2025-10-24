package strand

func ToRNA(dna string) string {
	dnaToRna := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	rna := []rune{}

	for _, ch := range dna {
		rna = append(rna, dnaToRna[ch])
	}

	return string(rna)
}
