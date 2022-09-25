package wordutil

import (
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {
	map_1 := make(map[string]int)

	v1 := strings.Fields(s)

	for _, word := range v1 {
		v2 := strings.ToLower(word)

		map_1[v2] = map_1[v2] + 1

	}
	return map_1
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
}
