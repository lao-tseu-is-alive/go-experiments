package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("💥 ERROR: %s expects a filename as an argument", os.Args[0])
	}
	filename := args[0]
	log.Printf("📂 Will try to read file: %s", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("💥 ERROR: trying to read file %s, error: %s", filename, err)
	}
	log.Printf("📂 File %s opened successfully", filename)
	log.Printf("📂 File content: \n%s", content)

}
