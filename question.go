package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Questions []Question

type Question struct {
	Question string `yaml:question`
	Answer   string `yaml:answer`
	Filename string `yaml:filename`
	Content  string
}

func loadQuestions() (Questions, error) {
	var questions Questions

	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return questions, err
	}

	err = yaml.Unmarshal(data, &questions)
	if err != nil {
		return questions, err
	}

	for i := 0; i < len(questions); i++ {
		err := questions[i].readFile()

		if err != nil {
			return questions, err
		}
	}

	return questions, nil
}

func (q *Question) readFile() error {
	data, err := ioutil.ReadFile("./files/" + q.Filename)

	if err != nil {
		return err
	}

	q.Content = string(data)

	return nil
}

func (qs *Questions) Loop() {
	for i := 0; i < len(*qs); i++ {
		question := (*qs)[i]

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
