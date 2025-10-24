package phonenumber

import (
	"fmt"
	"unicode"
	"errors"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	// clean up
	re := regexp.MustCompile(`[\(\)-\.\s]`)
	phoneNumber = string(re.ReplaceAll([]byte(phoneNumber), []byte("")))

	// check the length
	if len(phoneNumber) < 10 {
		return "", errors.New("not enough digits")
	}
	if len(phoneNumber) > 11 {
		return "", errors.New("too many digits")
	}

	// check for digits only
	for _, ch := range phoneNumber {
		if !unicode.IsDigit(ch) {
			return "", errors.New("non-digit characters and not allowed")
		}
	}

	// if 11 digits, check for valid country code
	if len(phoneNumber) == 11 && phoneNumber[0] != '1' {
		return "", errors.New("invalid country code")
	}

	// check area and exchange code
	var areaCodeStart int
	var exchangeCodeStart int
	if len(phoneNumber) == 11 {
		areaCodeStart = int(phoneNumber[1] - '0')
		exchangeCodeStart = int(phoneNumber[4] - '0')
	} else {
		areaCodeStart = int(phoneNumber[0] - '0')
		exchangeCodeStart = int(phoneNumber[3] - '0')
	}
	
	if areaCodeStart < 2 {
		return "", errors.New("invalid area code")
	}
	if exchangeCodeStart < 2 {
		return "", errors.New("invalid exchange code")
	}

	// omit country code
	if len(phoneNumber) == 11 {
		return phoneNumber[1:], nil
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	n, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%v) %v-%v", n[:3], n[3:6], n[6:]), nil
}
