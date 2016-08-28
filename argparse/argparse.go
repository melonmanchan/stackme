package argparse

import (
	"errors"
	"os"
	"strings"
)

// Reads arguments from stdin, returns them as a string
func ParseArguments() (output string, err error) {
	args := os.Args[1:]

	if len(args) == 0 {
		return "", errors.New("Query argument is missing!")
	}

	return strings.Join(args, " "), nil
}
