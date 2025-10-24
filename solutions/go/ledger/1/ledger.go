package ledger

import (
	"fmt"
	"slices"
	"time"
	"errors"
)

var ErrUnknownLocale = errors.New("unknown locale")
var ErrInvalidDate = errors.New("invalid date")
var ErrInvalidCurrency = errors.New("invalid currency")

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

var currencyChar = map[string]string{
	"EUR": "€",
	"USD": "$",
}

var locales = map[string]map[string]string{
	"en-US": map[string]string{
		"Date": "Date",
		"Description": "Description",
		"Change": "Change",
	},
	"nl-NL": map[string]string{
		"Date": "Datum",
		"Description": "Omschrijving",
		"Change": "Verandering",
	},
}

func formatChange(locale string, currency string, cents int) (string, error) {
	// check validity of currency.
	c, ok := currencyChar[currency]; if !ok {
		return "", ErrInvalidCurrency
	}

	// check for negative change.
	isNegative := false
	if cents < 0 {
		isNegative = true
		cents *= -1
	}

	// split the number into parts.
	thousands := 0
	units := 0

	if cents >= 1e5 {
		thousands = cents / 1e5
		cents %= 1e5
	}
	if cents >= 100 {
		units = cents / 100
		cents %= 100
	}

	s := ""

	switch locale {
	case "nl-NL":
		s += c + " "
		if thousands > 0 {
			s += fmt.Sprintf("%d.%03d,%02d", thousands, units, cents)
		} else if units > 0 {
			s += fmt.Sprintf("%d,%02d", units, cents)
		} else {
			s += fmt.Sprintf("0,%02d", cents)
		}
		if isNegative {
			s += "-"
		} else {
			s += " "
		}
	case "en-US":
		if isNegative {
			s += "("
		} else {
			s += " "
		}
		s += c
		if thousands > 0 {
			s += fmt.Sprintf("%d,%03d.%02d", thousands, units, cents)
		} else if units > 0 {
			s += fmt.Sprintf("%d.%02d", units, cents)
		} else {
			s += fmt.Sprintf("0.%02d", cents)
		}
		if isNegative {
			s += ")"
		} else {
			s += " "
		}
	default:
		return "", ErrUnknownLocale
	}

	return s, nil
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// copy entries
	var entriesCopy []Entry			
	for _, e := range entries {
		entriesCopy = append(entriesCopy, e)
	}

	// check currency
	if _, ok := currencyChar[currency]; !ok {
		return "", ErrInvalidCurrency
	}

	// sort by date, then by change.
	slices.SortFunc(entriesCopy, func(e1, e2 Entry) int {
		t1, _ := time.Parse("2006-01-02", e1.Date);
		t2, _ := time.Parse("2006-01-02", e2.Date);

		if t1 == t2 {
			return e1.Change - e2.Change
		}
		return int(t1.Sub(t2))
	})

	// building locale string.
	var s string

	// check locale
	l, ok := locales[locale]; if !ok {
		return "", ErrUnknownLocale
	}

	// format ledger header
	s = fmt.Sprintf("%-10s | %-25s | %s\n", l["Date"], l["Description"], l["Change"])

	for _, entry := range entriesCopy{
		// parse date
		date, err := time.Parse("2006-01-02", entry.Date); if err != nil {
			return "", err
		}

		// trimming description
		de := entry.Description
		if len(de) > 25 {
			de = de[:22] + "..."
		} else {
			de = fmt.Sprintf("%-25s", de)
		}

		// printing date in local format
		var d string
		if locale == "nl-NL" {
			d = date.Format("02-01-2006")
		} else if locale == "en-US" {
			d = date.Format("01/02/2006")
		}

		// formatting and printing change.
		a, err := formatChange(locale, currency, entry.Change)
		if err != nil {
			return "", err
		}

		// concat to the final string
		s += fmt.Sprintf("%-10s | %s | %13s\n", d, de, a)
	}

	return s, nil
}
