package main

import (
	"./argparse"
	"fmt"
	"log"
	"os"
)

func main() {
	query, err := argparse.ParseArguments(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(query)
}
