package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/adlerhsieh/interviewer/question"
	"gopkg.in/yaml.v2"
)

func readQuestions() []question.Question {
	var questions []question.Question

	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &questions)
	if err != nil {
		panic(err)
	}

	return questions
}

var questions = readQuestions()

func main() {
	for _, question := range questions {
		fmt.Println(question.Question)
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == question.Answer {
			fmt.Println("Correct\n")
		} else {
			fmt.Println("Incorrect\n")
		}
	}
}
