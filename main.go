package main

import (
	"log"
	"os"

	"./argparse"
	"./searcher"
)

func main() {
	query, err := argparse.ParseArguments(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	questions := []searcher.Question{}

	questions, err = searcher.GetStackOverflowQuestions(query)

	if err != nil {
		log.Fatal(err)
	}

	answers := []searcher.Answer{}

	answers, err = searcher.GetStackOverflowAnswers(questions)

	if err != nil {
		log.Fatal(err)
	}

	for _, a := range answers {
		log.Printf("%v", a)
	}
}
