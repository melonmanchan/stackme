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

	_, err = searcher.GetStackOverflowQuestions(query)

	if err != nil {
		log.Fatal(err)
	}
}
