package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RemoveAtIndex removes the element at the specified index from a string slice.
func RemoveAtIndex(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		fmt.Println("index out of range")
		return s
	}
	return append(s[:index], s[index+1:]...)
}

// RemoveIndexPlus2 removes the element at the specified index and the next element from a string slice.
func RemoveIndexPlus2(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		fmt.Println("index out of range")
		return s
	}
	return append(s[:index], s[index+2:]...)
}

// Uppercase converts elements to uppercase based on the "(up)" tag.
func Uppercase(s []string) []string {
	for i, val := range s {
		if val == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = RemoveAtIndex(s, i)
		}
	}
	return s
}

// Lowercase converts elements to lowercase based on the "(low)" tag.
func Lowercase(s []string) []string {
	for i, val := range s {
		if val == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = RemoveAtIndex(s, i)
		}
	}
	return s
}

// Capitalize capitalizes elements based on the "(cap)" tag.
func Capitalize(s []string) []string {
	for i, val := range s {
		if val == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = RemoveAtIndex(s, i)
		}
	}
	return s
}

// CaseN handles case changes with specified steps based on tags like "(up,2)".
func CaseN(s []string) []string {
	for idx, val := range s {
		if val == "(up," && idx < len(s)-1 {
			valMod := strings.TrimSuffix(s[idx+1], ")")
			number, _ := strconv.Atoi(valMod)
			for i := 1; i <= number; i++ {
				a := (idx - i)
				if a >= 0 && a < len(s) {
					s[a] = strings.ToUpper(s[a])
				}
			}
			s = RemoveIndexPlus2(s, idx)
		} else if val == "(low," {
			valMod := strings.TrimSuffix(s[idx+1], ")")
			number, _ := strconv.Atoi(string(valMod))
			for i := 1; i <= number; i++ {
				a := (idx - i)
				if a >= 0 && a < len(s) {
					s[a] = strings.ToLower(s[a])
				}
			}
			s = RemoveIndexPlus2(s, idx)
		} else if val == "(cap," {
			valMod := strings.TrimSuffix(s[idx+1], ")")
			number, _ := strconv.Atoi(string(valMod))
			for i := 1; i <= number; i++ {
				a := (idx - i)
				if a >= 0 && a < len(s) {
					s[a] = strings.Title(s[a])
				}
			}
			s = RemoveIndexPlus2(s, idx)
		}
	}
	return s
}

// HexToDec converts hexadecimal numbers in the slice to decimal numbers.
func HexToDec(s []string) []string {
	for i, v := range s {
		if v == "(hex)" {
			decVal, err := strconv.ParseInt(s[i-1], 16, 64)
			s[i-1] = strconv.Itoa(int(decVal))
			s = RemoveAtIndex(s, i)
			if err != nil {
				fmt.Printf("Error: %s is not a valid hexadecimal number\n", s)
			}
		}
	}
	return s
}

// BinToDec converts binary numbers in the slice to decimal numbers.
func BinToDec(s []string) []string {
	for i, v := range s {
		if v == "(bin)" {
			decVal, err := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = strconv.Itoa(int(decVal))
			s = RemoveAtIndex(s, i)
			if err != nil {
				fmt.Printf("Error: %s is not a valid binary number\n", s)
			}
		}
	}
	return s
}

// ArticleA corrects "a" to "an" based on the following word's initial letter.
func ArticleA(s []string) []string {
	for i, val := range s {
		if val == "a" && i > 0 && i+1 < len(s) {
			nextWord := s[i+1]
			vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
			for _, char := range vowels {
				if strings.HasPrefix(nextWord, char) {
					s[i] = val + "n"
				}
			}
		}
	}
	return s
}

// Punctuations handles punctuation marks by combining them with the preceding word.
func Punctuations(s []string) []string {
	spChars := []string{".", ",", "!", "?", ":", ";"}
	for idx, val := range s {
		for _, char := range spChars {
			if val == char {
				s[idx-1] = s[idx-1] + char
				s = RemoveAtIndex(s, idx)
			} else if strings.HasPrefix(val, char) {
				s[idx-1] = s[idx-1] + char
				s[idx] = s[idx][1:]
			}
		}
	}
	return s
}

// MultPuncts handles sequences of repeated identical punctuation marks.
func MultPuncts(s []string) []string {
	spChars := []string{",", "!", "?", ":", ";", "."}

	for i, val := range s {
		for _, v := range spChars {
			if len(val) > 1 && strings.HasPrefix(val, v) && strings.HasSuffix(val, v) {
				s[i-1] = s[i-1] + val
				s = RemoveAtIndex(s, i)
			}
		}
	}
	return Punctuations(s)
}

// Apostrophe manages apostrophes by combining them with the following or preceding word.
func Apostrophe(s []string) []string {
	count := 0
	for i, v := range s {
		if v == "'" && count == 0 {
			s[i+1] = v + s[i+1]
			s = append(s[:i], s[i+1:]...)
			count++
		}
	}

	for i, w := range s {
		if w == "'" {
			s[i-1] += w
			count = 0
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	args := os.Args[1:] // Get command-line arguments excluding the program name

	// Read the content of the file specified in the first command-line argument
	text, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}

	// Split the content of the file into a slice of words
	strSlice := strings.Split(string(text), " ")

	// Apply a series of transformations to the word slice
	strSlice = CaseN(strSlice)
	strSlice = Uppercase(strSlice)
	strSlice = Lowercase(strSlice)
	strSlice = Capitalize(strSlice)
	strSlice = Punctuations(strSlice)
	strSlice = MultPuncts(strSlice)
	strSlice = HexToDec(strSlice)
	strSlice = BinToDec(strSlice)
	strSlice = ArticleA(strSlice)
	strSlice = Apostrophe(strSlice)

	// Join the modified word slice back into a string
	var sentence string
	sentence = strings.Join(strSlice, " ") + "\n"

	// Create a new file named "result.txt" and write the modified string into it
	file, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the modified string to the file
	_, err = file.WriteString(sentence)
	if err != nil {
		panic(err)
	}
}
