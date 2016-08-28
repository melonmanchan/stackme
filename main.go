package main

import (
	"./argparse"
	"./searcher"
	"log"
	"os"
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
