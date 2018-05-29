package main

import (
	"log"
	"os"
)

func main() {
	questions, err := loadQuestions()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	questions.Loop()
}
