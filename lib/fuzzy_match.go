package lib

import (
	"strings"
)

type MatchResult struct {
	StringMatches []string
}

func FuzzyMatches(reference []string, searchTerm string) *MatchResult {
	m := &MatchResult{[]string{}}
	for _, item := range reference {
		if consonantMatch(searchTerm, item) {
			m.StringMatches = append(m.StringMatches, item)
		} // end if
	} // end for loop
	return m
}

func transformToConsonants(testString string) string {
	var result []string
	chars := strings.Split(testString, "")
	for _, char := range chars {
		if !strings.Contains("aeiou", char) {
			result = append(result, char)
		}
	}

	return strings.Join(result, "")
}

func consonantMatch(input, testString string) (result bool) {
	if len(input) > len(testString) {
		return
	} else if input == transformToConsonants(testString) {
		return true
	} else {
		return
	}
}

