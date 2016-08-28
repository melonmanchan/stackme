package searcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Question struct {
	IdAnswer int    `json:accepted_answer_id`
	Title    string `json:title`
	Link     string `json:link`
}

type Questions struct {
	Questions []Question `json:items`
}

func URLEncodeString(str string) (string, error) {
	encoded, err := url.Parse(str)

	if err != nil {
		return "", err
	}

	return encoded.String(), nil
}

func FormatSearchToURL(question string) (output string) {
	formatted, _ := URLEncodeString(question)

	return "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=relevance&accepted=True&site=stackoverflow&q=" + formatted
}

func SearchByQuery(query string) (io.ReadCloser, error) {

	url := FormatSearchToURL(query)

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func GetStackOverflowQuestions(question string) (string, error) {
	body, err := SearchByQuery(question)

	if err != nil {
		return "", err
	}

	var questions = new(Questions)

	err = json.NewDecoder(body).Decode(questions)
	body.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%d", len(questions.Questions))
	}

	return "", nil
}
