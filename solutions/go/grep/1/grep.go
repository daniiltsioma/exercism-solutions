package grep

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func Search(pattern string, flags, files []string) []string {
	// handle flags.
	prependLineNumbers := false
	listFilenames := false
	caseInsensitive := false
	invertMatches := false
	matchEntireLines := false
	for _, f := range flags {
		switch f {
		case "-n":
			prependLineNumbers = true
		case "-l":
			listFilenames = true
		case "-i":
			caseInsensitive = true
		case "-v":
			invertMatches = true
		case "-x":
			matchEntireLines = true
		}
	}

	// convert pattern to lowercase
	if caseInsensitive {
		pattern = strings.ToLower(pattern)
	}

	results := []string{}

	matchedFilenames := map[string]bool{}

	for _, filename := range files {
		file, err := os.Open(filename)		
		if err != nil {
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		lineNumber := 0
		for scanner.Scan() {
			isMatch := false
			originalLine := scanner.Text()
			line := string([]byte(originalLine))
			lineNumber++

			// case sensitive
			if caseInsensitive {
				line = strings.ToLower(line)
			}

			if matchEntireLines {
				isMatch = line == pattern
			} else {
				isMatch = strings.Contains(line, pattern)
			}

			if invertMatches {
				isMatch = !isMatch
			}

			if !isMatch {
				continue
			}

			if listFilenames {
				matchedFilenames[filename] = true
				continue
			}

			res := ""
			
			if len(files) > 1 {
				res += filename + ":"
			}
			if prependLineNumbers {
				res += strconv.Itoa(lineNumber) + ":"
			}
			res += originalLine
			results = append(results, res)
		}
	}

	if listFilenames {
		for fn := range matchedFilenames {
			results = append(results, fn)
		}
	}

	return results
}
