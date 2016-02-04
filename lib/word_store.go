package lib

import (
	"io/ioutil"
	"strings"
)

const WORD_FILE = "/data/words.txt"

func WordStore() []string { // need to initialize this in the `init` function for this package in main
	contents, err := ioutil.ReadFile(WORD_FILE)
	if err != nil {
		panic(err)
	}
	text := string(contents)
	words := strings.Split(text, "\n")

	return words
}
