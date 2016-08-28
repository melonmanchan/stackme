package searcher

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func URLEncodeString(str string) (string, error) {
	encoded, err := url.Parse(str)

	if err != nil {
		return "", err
	}

	return encoded.String(), nil
}

func FormatSearchToURL(question string) (output string) {
	formatted, _ := URLEncodeString(question)

	return "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=votes&accepted=True&site=stackoverflow&q=" + formatted
}

func SearchByQuery(query string) (string, error) {

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
