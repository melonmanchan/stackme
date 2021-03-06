package searcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DoHTTPGet performs a HTTP get request
func DoHTTPGet(url string) (io.ReadCloser, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

// QuestionsIDToColonSeparateString creates a semicolon-separated string of IDs
func QuestionsIDToColonSeparateString(questions []Question) (string, error) {
	if len(questions) == 0 {
		return "", errors.New("Length of questions cannot be 0!")
	}

	var output string

	for _, question := range questions {
		output += strconv.Itoa(question.QuestionID) + ";"
	}

	// Remove the last semicolon from output
	var trimmed = strings.TrimSuffix(output, ";")

	return trimmed, nil
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

// FormatAnswersURL gets an url that fetches an array of answers for questions
func FormatAnswersURL(questions []Question) (string, error) {
	formatted, err := QuestionsIDToColonSeparateString(questions)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://api.stackexchange.com/2.2/questions/%s/answers?filter=withbody&site=stackoverflow", formatted), nil
}

// GetQuestionResponse searches Stack Overflow questions by a question
func GetQuestionResponse(query string) (io.ReadCloser, error) {
	url := FormatSearchToURL(query)

	response, err := DoHTTPGet(url)

	return response, err
}

// GetAnswerResponse gets
func GetAnswerResponse(questions []Question) (io.ReadCloser, error) {
	var response io.ReadCloser

	url, err := FormatAnswersURL(questions)

	log.Printf(url)

	if err != nil {
		return nil, err
	}

	response, err = DoHTTPGet(url)

	return response, err
}

// GetStackOverflowQuestions Fetches a list of questions from stackoverflow matching a query
func GetStackOverflowQuestions(question string) ([]Question, error) {
	body, err := GetQuestionResponse(question)

	if err != nil {
		return nil, err
	}

	var questions = new(Questions)

	err = json.NewDecoder(body).Decode(questions)
	body.Close()

	if err != nil {
		return nil, err
	}

	return questions.Items, nil
}

// GetStackOverflowAnswers returns an array of answers of questions
func GetStackOverflowAnswers(questions []Question) ([]Answer, error) {
	body, err := GetAnswerResponse(questions)

	if err != nil {
		return nil, err
	}

	var answers = new(Answers)

	err = json.NewDecoder(body).Decode(answers)
	body.Close()

	if err != nil {
		return nil, err
	}

	return answers.Items, nil
}
