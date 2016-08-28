package argparse_test

import (
	"./"
	"fmt"
	"testing"
)

var ok = []string{"stackme", "this is a working test"}
var bad = []string{"this is bad input!"}
var nonexistant = []string{}

func TestArgParse(t *testing.T) {
	output, err := argparse.ParseArguments(ok)

	if err != nil {
		t.Error(err)
	}

	if output != ok[1] {
		t.Error(fmt.Sprintf("Output did not match input! Got %s but expected %s", output, ok[1]))
	}

	_, err = argparse.ParseArguments(bad)

	if err == nil {
		t.Error("ParseArguments should have errored with no parameters!")
	}

	output, err = argparse.ParseArguments(nonexistant)

	if output != "" && err != nil {
		t.Error("ParseArguments should have errored with an empty array!")
	}
}
