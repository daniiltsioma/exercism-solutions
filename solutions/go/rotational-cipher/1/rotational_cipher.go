package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {	
	rotated := []rune{}

	for _, r := range plain {
		if !unicode.IsLetter(r) {
			rotated = append(rotated, r)
			continue
		}
		shifted := int(r) + shiftKey
		if unicode.IsUpper(r) {
			if shifted > 90 {
				rotated = append(rotated, rune(64 + shifted - 90))
			} else {
				rotated = append(rotated, rune(shifted))
			}
		} else {
			if shifted > 122 {
				rotated = append(rotated, rune(96 + shifted - 122))
			} else {
				rotated = append(rotated, rune(shifted))
			}
		}
	}

	return string(rotated)
}
