package lib

import (
	"sync"
)

func BuildSearchResults(searchTerm string) []string {
	var builtResults []string
	sectionDepth := len(searchTerm) / 2
	wg := &sync.WaitGroup{}
	for i := 1; i <= sectionDepth; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			// might need to add a mutex lock
			builtResults = append(builtResults, NewSectionedString(searchTerm, i).SearchResults()...)
		}()

		wg.Wait()
	}

	return builtResults
}
