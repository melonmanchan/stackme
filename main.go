package main

import (
	"./argparse"
	"./searcher"
	"fmt"
	"log"
	"os"
)

func main() {
	query, err := argparse.ParseArguments(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	results, err := searcher.SearchByQuery(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(results)
}
