package transform

import (
	delete "go-reloaded/deletion" // Importing the deletion package
	"strings"
)

// Punctuation merges adjacent punctuation marks or removes single ones
func Punctuation(s []string) []string {
	// Define a slice containing common punctuation marks
	spChars := []string{".", ",", "!", "?", ":", ";"}

	for i, v := range s {
		for _, c := range spChars {
			// If the current element matches a punctuation mark
			if v == c {
				// Merge it with the previous element
				s[i-1] += s[i]
				// Remove the current element
				s = delete.RemoveAtIndex(s, i)
			} else if strings.HasPrefix(v, c) {
				// If the current element starts with a punctuation mark
				// Merge the punctuation mark with the previous element
				s[i-1] += c
				// Remove the punctuation mark from the current element
				s[i] = s[i][1:]
				// Move the current element to the previous index
				s = append(s[:i], s[i:]...)
			}
		}
	}
	return s
}

// MultPunctuation handles multiple adjacent punctuation marks
func MultPunctuation(s []string) []string {
	// Define a slice containing common punctuation marks
	spChars := []string{".", ",", "!", "?", ":", ";"}

	for i, v := range s {
		for _, c := range spChars {
			// If the current element starts and ends with the same punctuation mark and its length is greater than 1
			if strings.HasPrefix(v, c) && strings.HasSuffix(v, c) && len(v) > 1 {
				// Merge it with the previous element
				s[i-1] += s[i]
				// Remove the current element
				s = delete.RemoveAtIndex(s, i)
			}
		}
	}
	// Call Punctuation function to handle merged punctuation marks
	return Punctuation(s)
}

// Apostrophe handles placement of apostrophes
func Apostrophe(s []string) []string {
	count := 0
	// Loop through the elements to find the first occurrence of an apostrophe
	for i, v := range s {
		if v == "'" && count == 0 {
			// Merge it with the next element
			s[i+1] = v + s[i+1]
			// Remove the apostrophe element
			s = append(s[:i], s[i+1:]...)
			count++
		}
	}
	// Loop through the elements again to find any remaining apostrophes
	for i, w := range s {
		if w == "'" {
			// Merge it with the previous element
			s[i-1] += w
			// Reset the count
			count = 0
			// Remove the apostrophe element
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// ArticleA handles placement of the article "a"
func ArticleA(s []string) []string {
	vowelsLow := []string{"a", "e", "i", "o", "u", "h"}
	vowelsCap := []string{"A", "E", "I", "O", "U", "H"}

	for i, v := range s {
		// If the current element is "a"
		if v == "a" {
			for _, vwl := range vowelsLow {
				// If the next element starts with a vowel or 'h', add "n" to the current element
				if strings.HasPrefix(s[i+1], vwl) {
					s[i] += "n"
				}
			}
		} else if v == "A" {
			for _, vwl := range vowelsCap {
				// If the next element starts with a vowel or 'H', add "N" to the current element
				if strings.HasPrefix(s[i+1], vwl) {
					s[i] += "N"
				}
			}
		}

		// If the current element is "A"
		if v == "A" {
			for _, vwl := range vowelsLow {
				// If the next element starts with a vowel or 'h', add "n" to the current element
				if strings.HasPrefix(s[i+1], vwl) {
					s[i] += "n"
				}
			}
		}
	}
	return s
}
