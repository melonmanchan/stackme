package argparse

import (
	"errors"
	"strings"
)

// Reads arguments from stdin, returns them as a string
func ParseArguments(input []string) (output string, err error) {
	args := input[1:]

	if len(args) == 0 {
		return "", errors.New("Query argument is missing!")
	}

	return strings.Join(args, " "), nil
}
