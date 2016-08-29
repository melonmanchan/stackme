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

	err = searcher.GetStackOverflowAnswers(questions)
}
