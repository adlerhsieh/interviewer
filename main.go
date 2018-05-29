package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
var hasError bool

func main() {
	for i := 0; i < len(questions); i++ {
		q := &questions[i]

		data, err := ioutil.ReadFile("./question/files/" + q.Filename)
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				fmt.Println("File not found:\n  question: " + q.Question + "\n  file location: ./question/files/" + q.Filename)
				hasError = true
			} else {
				panic(err)
			}
		}

		q.Content = string(data)
	}

	if hasError {
		os.Exit(1)
	}

	for _, question := range questions {
		fmt.Println(question.Content)
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
