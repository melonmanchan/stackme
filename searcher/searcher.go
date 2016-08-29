package searcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Question is a datatype representing a single question from SO
type Question struct {
	Title            string
	Link             string
	AcceptedAnswerID int `json:"Accepted_answer_id"`
	QuestionID       int `json:"question_id"`
}

// Questions is a  datatype containing a nested array of questions
type Questions struct {
	Items []Question
}

// URLEncodeString  encodes a string as URL-safe
func URLEncodeString(str string) (string, error) {
	encoded, err := url.Parse(str)

	if err != nil {
		return "", err
	}

	return encoded.String(), nil
}

// FormatSearchToURL appends a question to a SO API url
func FormatSearchToURL(question string) (output string) {
	formatted, _ := URLEncodeString(question)

	return "https://api.stackexchange.com/2.2/search/advanced?order=desc&sort=relevance&accepted=True&site=stackoverflow&q=" + formatted
}

// SearchByQuery does a GET request to the StackOverflow API, returning it's body
func SearchByQuery(query string) (io.ReadCloser, error) {

	url := FormatSearchToURL(query)

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// GetStackOverflowQuestions Fetches a list of questions from stackoverflow matching a query
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
		return "", err
	}

	for _, item := range questions.Items {
		fmt.Println(item.Title)
		fmt.Println(item.AcceptedAnswerID)
	}

	return "", nil
}
