package wordutil

import (
	"strings"
)

// Finds first occurrence of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// the index of the first occurence of the word in the input string as the
// corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndex(s string) map[string]int {
	map1 := make(map[string]int)

	v1 := strings.Fields(s)

	for _, word := range v1 {
		v2 := strings.ToLower(word)
		_, ok := map1[v2]
		if !ok {
			map1[v2] = map1[v2] + strings.Index(s, word)
		}
	}

	return map1

	// TODO: implement me
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful
}
