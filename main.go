package main

import (
	"log"
	"os"
	"unicode"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("💥 ERROR: %s expects a text filename as first argument", os.Args[0])
	}
	filename := args[0]
	log.Printf("📂 Will try to read file: %s", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("💥 ERROR: trying to read file %s, error: %s", filename, err)
	}
	log.Printf("📂 File %s opened successfully", filename)
	numBytes := len(content)
	// count the number of lines and words in the content
	numLines, numWords := 0, 0
	inWord := false
	for _, rune := range string(content) {
		if rune == '\n' {
			numLines++
		}
		if unicode.IsSpace(rune) {
			if inWord {
				inWord = false // Exiting a word
			}
		} else {
			if !inWord {
				numWords++
				inWord = true // Entering a word
			}
		}
	}
	log.Printf("📂 File num lines, words and bytes: %d\t%d\t%d", numLines, numWords, numBytes)

}
