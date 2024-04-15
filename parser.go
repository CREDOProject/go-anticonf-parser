package goanticonfparser

import (
	"regexp"
	"strings"
)

var variableMatcher = regexp.MustCompile(`(?m)^(?:[a-zA-Z][[:alnum:]]*_?)(?:[[:alnum:]]*_?)+=".+"$`)

// Parse variables present in a configure file.
// Returns a map[string]string of all variables.
func Parse(file string) map[string]string {
	allStrings := variableMatcher.FindAllString(file, -1)
	var kv = make(map[string]string)
	for _, line := range allStrings {
		s := strings.Split(line, "=")
		if len(s) > 1 {
			key := s[0]
			val := strings.Trim(s[1], "\"")
			kv[key] = val
		}
	}
	return kv
}
