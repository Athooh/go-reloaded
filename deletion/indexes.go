package delete

import (
	"fmt"
)

func RemoveAtIndex(s []string, i int) []string {
	if i <= 0 || i > len(s) {
		fmt.Println("Index Out of range.")
	} else {
		s = append(s[:i], s[i+1:]...)
	}
	return s
}

func RemoveAtIndex2(s []string, i int) []string {
	if i <= 0 || i > len(s) {
		fmt.Println("Index Out of range.")
	} else {
		s = append(s[:i], s[i+2:]...)
	}
	return s
}
