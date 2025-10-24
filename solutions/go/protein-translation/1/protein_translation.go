package protein

import "errors"

var ErrStop = errors.New("STOP keyword entered")
var ErrInvalidBase = errors.New("Invalid codon base")

func FromRNA(rna string) ([]string, error) {
	acids := []string{}

	currentCodon := []rune{}
	for _, ch := range rna {
		currentCodon = append(currentCodon, ch)
		if len(currentCodon) == 3 {
			acid, err := FromCodon(string(currentCodon))
			if err == ErrStop {
				return acids, nil
			} else if err == ErrInvalidBase {
				return nil, ErrInvalidBase
			} else {
				acids = append(acids, acid)
				currentCodon = currentCodon[:0]
			}
		}
	}

	return acids, nil
}

func FromCodon(codon string) (string, error) {
	codonToRNA := map[string]string {
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	v, ok := codonToRNA[codon]; if !ok {
		return "", ErrInvalidBase
	} else if v == "STOP" {
		return "", ErrStop
	} else {
		return v, nil
	}
}