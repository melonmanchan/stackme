package searcher_test

import (
	"."
	"testing"
)

func TestURLEncode(t *testing.T) {
	encoded, err := searcher.URLEncodeString("these spaces should be encoded")
	if err != nil {
		t.Error(err)
	}

	if encoded != "these%20spaces%20should%20be%20encoded" {
		t.Error("Error encoding spaces: got: " + encoded)
	}
}
