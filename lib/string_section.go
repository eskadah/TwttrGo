// given a string and template length, this file creates a struct with all the string sections
package lib

import (
	"strings"
)

var Cache = WordStore()

const MIN_TEMPLATE_LENGTH = 2

type SectionedString struct {
	RootString     string
	templateLength int
	templates      [][]int
	Sections       [][]string
}

type SearchResult []string

func NewSectionedString(root string, templateLength int) *SectionedString {
	templates := defineTemplates(len(root), templateLength)
	self := &SectionedString{root, templateLength, templates, [][]string{}}
	self.generateSections()
	return self
}

func defineTemplates(rootLength, templateLength int) (templates [][]int) {
	if templateLength == MIN_TEMPLATE_LENGTH {
		for unitLength := MIN_TEMPLATE_LENGTH; unitLength <= rootLength - MIN_TEMPLATE_LENGTH; unitLength++ {
			templates = append(templates, []int{unitLength, rootLength - unitLength})
		}
	} else {
		for unitLength := MIN_TEMPLATE_LENGTH; unitLength <= rootLength - MIN_TEMPLATE_LENGTH; unitLength++ {
			for _, template := range defineTemplates(unitLength, templateLength - 1) {
				template = append(template, rootLength - unitLength)
				templates = append(templates, template)
			}
		}
	}
	return
}

func (sectionedStr *SectionedString) generateSection(template []int) {
	var result []string
	root := sectionedStr.RootString
	start_index := 0
	for _, unitLength := range template {
		final_index := start_index + unitLength
		result = append(result, root[start_index:final_index])
		start_index += unitLength
	}
	sectionedStr.Sections = append(sectionedStr.Sections, result)
}

func (sectionedStr *SectionedString) generateSections() {
	for _, template := range sectionedStr.templates {
		sectionedStr.generateSection(template)
	}
}

// add case for no sections and templateLength == 1
func (section *SectionedString) matches() [][]*MatchResult {
	matches := [][]*MatchResult{}
	if section.templateLength == 1 {
		matches = append(matches, []*MatchResult{FuzzyMatches(Cache, section.RootString)})
	} else {
		for _, section := range section.Sections {
			sectionMatch := []*MatchResult{}
			for _, unit := range section {
				unitMatch := FuzzyMatches(Cache, unit)
				sectionMatch = append(sectionMatch, unitMatch)
			}
			matches = append(matches, sectionMatch)
		}
	}
	return matches
}

func (section *SectionedString) SearchResults() (searchResults SearchResult) {
	for _, matchResults := range section.matches() {
		searchResults = append(searchResults, multiplex(mapToString(matchResults)...)...)
	}
	return
}

func mapToString(matches []*MatchResult) (strings [][]string) {
	for _, match := range matches {
		strings = append(strings, match.StringMatches)
	}
	return
}

func multiplex(arrays ...[]string) (result []string) {
	if len(arrays) == 1 {
		result = append(result, strings.Join(arrays[0], " "))
	} else if len(arrays) == 2 {
		first, second := arrays[0], arrays[1]
		for _, firstUnit := range first {
			for _, secondUnit := range second {

				result = append(result, strings.Join([]string{firstUnit, secondUnit}, " "))
			}
		}
	} else if len(arrays) > 2 {
		result = multiplex(arrays[0], arrays[1])
		for i := 2; i < len(arrays); i++ {
			result = multiplex(result, arrays[i])

		}
	}
	return
}
