# TwttrGo
A go lexer/parser for creating sentences from consonants. It has a single JSON API endpoint that delivers data to a Backbone.js SPA

The application parses consonant strings and constructs single word matches and sentences based on search results from a specified data file.

The rules for parsing are:

* Each word must contain at least 2 consonants and any number of vowels.
* There can only be the specified number and arrangement of input consonants in the parsed results
* All possible arrangements of parsed strings should be used to construct sentences in the input arrangement.
