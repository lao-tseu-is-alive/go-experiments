package main

import (
	"log"
	"os"
	"unicode"
)

const (
	DEBUG bool = false
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
	numLines, numWords, numLettersInCurrentWord := 1, 0, 0
	inWord := false
	for index, rune := range string(content) {
		if unicode.IsSpace(rune) || unicode.IsPunct(rune) || !unicode.IsLetter(rune) {
			if inWord {
				inWord = false // Exiting a word
				if numLettersInCurrentWord > 1 {
					numWords++
				}
				numLettersInCurrentWord = 0
			}
		} else {
			if unicode.IsLetter(rune) && !inWord {
				// Entering a word
				inWord = true
				numLettersInCurrentWord++
			} else {
				// Inside a word
				numLettersInCurrentWord++
			}
		}

		if DEBUG {
			if unicode.IsLetter(rune) {
				log.Printf("👍 @ byte %d, rune: \t%c \tis a Letter. inWord=%t numLettersInCurrentWord=%d \tnumWords: %d", index, rune, inWord, numLettersInCurrentWord, numWords)
			} else {
				log.Printf("❗ @ byte %d, rune: \t%c \tis NOT a Letter. inWord=%t numLettersInCurrentWord=%d \tnumWords: %d", index, rune, inWord, numLettersInCurrentWord, numWords)
			}
		}
		if rune == '\n' {
			if numWords != numLines*4 {
				log.Printf("❌ At line \t%d, found \t%d words (should be %d)", numLines, numWords, numLines*4)
			} else {
				log.Printf("✅ At line \t%d, found \t%d words", numLines, numWords)
			}
			numLines++
		}

	}
	log.Printf("📂 File num lines, words and bytes: %d\t%d\t%d", numLines, numWords, numBytes)

}
