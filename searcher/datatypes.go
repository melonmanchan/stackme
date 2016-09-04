package searcher

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

// Answer is a datatype representing a single answer from SO
type Answer struct {
	Body     string
	Score    int
	AnswerID int `json:"answer_id"`
}

// Answers is a datatype containing a nested array of answers
type Answers struct {
	Answers []Answer
}
