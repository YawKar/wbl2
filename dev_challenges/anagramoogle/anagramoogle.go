package anagramoogle

import (
	"fmt"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Input: slice of utf8-encoded russian words
//
// Output: pointer to map of anagram sets,
// where key is the first word appeared in the input,
// value is a sorted (strings.Compare) unique slice of all words in the set
//
// Note: anagram sets with power of 1 won't be in the output;
// all words will be lower-cased
func GroupIntoAnagrams(words []string) (*map[string][]string, error) {
	for _, word := range words {
		if !utf8.ValidString(word) {
			return nil, fmt.Errorf("Word %q is not a valid utf8 string", word)
		}
		if strings.ContainsFunc(word, func(r rune) bool {
			return !unicode.Is(unicode.Cyrillic, r) || !unicode.IsLetter(r)
		}) {
			return nil, fmt.Errorf("Word %q contains non-cyrillic or forbidden symbol", word)
		}
	}
	sortedToSlices := make(map[string][]string)

	words = slices.Clone(words)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	for _, word := range words {
		wordK := toSortedKey(word)
		slice, found := sortedToSlices[wordK]
		if !found {
			slice = make([]string, 0, 1)
		}
		slice = append(slice, word)
		sortedToSlices[wordK] = slice
	}
	result := make(map[string][]string)
	for _, v := range sortedToSlices {
		firstAppearance := v[0]
		if len(v) == 1 {
			continue // skip sets with the power of 1
		}
		slices.Sort(v)
		v = keepUnique(v)
		result[firstAppearance] = v
	}
	return &result, nil
}

// modifies in-place, `s` should be sorted
func keepUnique[T comparable](s []T) []T {
	curx := 1
	setx := 1
	for curx < len(s) {
		if s[curx] != s[curx-1] {
			s[setx] = s[curx]
			setx++
		}
		curx++
	}
	return s[:setx]
}

func toSortedKey(word string) string {
	runed := []rune(strings.ToLower(word))
	slices.Sort(runed)
	return string(runed)
}
