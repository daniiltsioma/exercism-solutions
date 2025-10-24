package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~\*=-]*>`)	
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"[^"]*"`)

	quotedSubstrings := []string{}
	for _, l := range lines {
		quotedSubstrings = append(quotedSubstrings, re.FindStringSubmatch(l)...)
	}

	count := 0
	for _, s := range quotedSubstrings {
		re = regexp.MustCompile(`(?i)password`)
		if re.MatchString(s) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	text = re.ReplaceAllString(text, "")
	return text
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User[\s]+([\w]+)`)	

	for i, l := range lines {
		if re.MatchString(l) {
			match := re.FindStringSubmatch(l)
			lines[i] = fmt.Sprintf("[USR] %s %s", match[1], l)
		}
	}

	return lines
}
