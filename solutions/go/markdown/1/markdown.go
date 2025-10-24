package markdown

import (
	"fmt"
	"strings"
	"regexp"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	// parse bold and italics
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	listOpened := false
	html := ""

	lines := strings.Split(markdown, "\n")

	// match header markdown.
	h := regexp.MustCompile(`#+\s`)

	for _, line := range lines {

		// headers
		// check if line starts with pound (#) characters
		loc := h.FindIndex([]byte(line)); if loc != nil && loc[0] == 0 {
			// determine the level
			lvl := loc[1] - loc[0] - 1
			if lvl < 7 { // valid heading
				line = fmt.Sprintf("<h%d>", lvl) + string(line[loc[1]:]) + fmt.Sprintf("</h%d>", lvl)
			} else {
				line = "<p>" + line + "</p>"
			}
			html += line
			continue
		}

		// unordered items
		if line[:2] == "* " {
			if !listOpened {
				html += "<ul>"
				listOpened = true
			}
			html += "<li>" + line[2:] + "</li>"
			continue
		} else if listOpened {
			html += "</ul>"
			listOpened = false
		}

		// paragraph
		line = "<p>" + line + "</p>"
		html += line
	}

	if listOpened {
		html += "</ul>"
		listOpened = false
	}

	return html
}
