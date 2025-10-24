package cipher

import (
	"unicode"
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift struct {
	distance int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return shift{ distance: 3}
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift{ distance: distance }
}

func normalizeLowercaseAscii(ascii int) int {
	if ascii < 97 {
		return 123 - (97 - ascii)
	}
	if ascii > 122 {
		return 96 + (ascii - 122)
	}
	return ascii
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	res := []rune{}

	for _, ch := range input {
		if !unicode.IsLetter(ch) {
			continue
		}
		ascii := int(ch) + c.distance
		ascii = normalizeLowercaseAscii(ascii)
		res = append(res, rune(ascii))
	}

	return string(res)
}

func (c shift) Decode(input string) string {
	input = strings.ToLower(input)
	res := []rune{}

	for _, ch := range input {
		ascii := int(ch) - c.distance
		ascii = normalizeLowercaseAscii(ascii)
		res = append(res, rune(ascii))
	}

	return string(res)
}

func NewVigenere(key string) Cipher {
	// test that key is valid
	valid := false
	for _, ch := range key {
		if int(ch) < 97 || int(ch) > 122 {
			return nil
		}
		if ch != 'a' {
			valid = true
		}	
	}
	if !valid {
		return nil
	}

	return vigenere{ key: key }
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	res := []rune{}

	for _, ch := range input {
		if !unicode.IsLetter(ch) {
			continue
		}
		ascii := int(ch) + int(v.key[len(res) % len(v.key)]) - 97
		ascii = normalizeLowercaseAscii(ascii)
		res = append(res, rune(ascii))
	}

	return string(res)
}

func (v vigenere) Decode(input string) string {
	res := []rune{}	

	for _, ch := range input {
		ascii := int(ch) - (int(v.key[len(res) % len(v.key)]) - 97)
		ascii = normalizeLowercaseAscii(ascii)
		res = append(res, rune(ascii))
	}

	return string(res)
}
