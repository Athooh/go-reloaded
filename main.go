package main

import (
	"fmt"                                  // Importing the fmt package for printing messages
	transform "go-reloaded/transformation" // Importing a custom package for text transformation
	"os"                                   // Importing the os package for working with operating system functionality
	"strings"                              // Importing the strings package for string manipulation
)

func main() {
	// Checking if the number of command line arguments is not equal to 3
	if len(os.Args) != 3 {
		fmt.Println("Error! Invalid Input") // Printing an error message if the number of arguments is invalid
		return                              // Exiting the program
	}

	args := os.Args[1:] // Extracting command line arguments, excluding the program name

	// Reading the content of the input file specified in the first command line argument
	inputFile, err := os.ReadFile(args[0])
	if err != nil { // Checking if there was an error while reading the file
		fmt.Println("Error!", err) // Printing the error message
		return                     // Exiting the program
	}

	// Splitting the content of the input file into a slice of words
	slice := strings.Split(string(inputFile), " ")

	// Performing various text transformations on the slice of words
	slice = transform.Lowercase(slice)
	slice = transform.Uppercase(slice)
	slice = transform.Capitalize(slice)
	slice = transform.HexToDec(slice)
	slice = transform.BinToDec(slice)
	slice = transform.Punctuation(slice)
	slice = transform.MultPunctuation(slice)
	slice = transform.Apostrophe(slice)
	slice = transform.ArticleA(slice)

	// Joining the transformed words back into a single text string
	text := strings.Join(slice, " ")

	// Creating an output file specified in the second command line argument
	outputFile, err := os.Create(args[1])
	if err != nil { // Checking if there was an error while creating the file
		fmt.Println("Error!", err) // Printing the error message
		return                     // Exiting the program
	}

	// Writing the transformed text to the output file
	_, err = outputFile.WriteString(text)
	if err != nil { // Checking if there was an error while writing to the file
		fmt.Println("Error!", err) // Printing the error message
		return                     // Exiting the program
	}
}
