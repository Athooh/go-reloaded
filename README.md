*Text Editor Application*

This project is all about building a versatile text processing tool in Go. The tool will read an input text file, apply specific modifications, and then save the modified text into another file.
Project Objectives

For this project, I'll be leveraging some of my existing functions from my previous repository to implement text completion and editing functionalities. The main objectives include:

Implementing various modifications such as converting hexadecimal and binary numbers, transforming text case (uppercase, lowercase, capitalized), and formatting punctuation correctly.
Ensuring the code adheres to best practices and includes unit tests for robustness.

Features and Modifications

Here's what my text editor application will do:

    Conversion of Hexadecimal and Binary Numbers:
        Replace (hex) with the decimal version of the preceding hexadecimal number.
        Replace (bin) with the decimal version of the preceding binary number.

    Case Transformations:
        (up) converts the preceding word to uppercase.
        (low) converts the preceding word to lowercase.
        (cap) converts the preceding word to capitalized form.
        For (low, <number>), (up, <number>), and (cap, <number>), apply case transformation to the specified number of preceding words.

    Punctuation Formatting:
        Ensure proper spacing around common punctuation marks (.,!?;:).
        Handle special cases for groups of punctuation (e.g., ..., !?).

    Handling Apostrophes:
        Properly position apostrophes (') around words or phrases.

    Indefinite Article Replacement:
        Replace a with an before words starting with a vowel or 'h'.

Usage:

$ cat sample.txt"

it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt

It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.

$ cat sample.txt

Simply add 42 (hex) and 10 (bin) and you will see the result is 68.

$ go run . sample.txt result.txt

$ cat result.txt

Simply add 66 and 2 and you will see the result is 68.

$ cat sample.txt

There is no greater agony than bearing a untold story inside you.

$ go run . sample.txt result.txt

$ cat result.txt

There is no greater agony than bearing an untold story inside you.

$ cat sample.txt

Punctuation tests are ... kinda boring ,don't you think !?

$ go run . sample.txt result.txt

$ cat result.txt

Punctuation tests are... kinda boring, don't you think!?

This project will help you learn about :

    The Go file system(fs) API
    String and numbers manipulation
