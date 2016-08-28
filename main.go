package main

import (
	"./argparse"
	"fmt"
	"log"
)

func main() {
	query, err := argparse.ParseArguments()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(query)
}
