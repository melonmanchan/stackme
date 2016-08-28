package searcher

import (
	"bytes"
	"fmt"
	"net/http"
)

func FormatSearchToURL(question string) (output string) {
	return "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=votes&accepted=True&site=stackoverflow&q=" + question
}

func SearchByQuery(query string) (output string, err error) {

	url := FormatSearchToURL(query)

	fmt.Println(url)

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	return buf.String(), nil
}
