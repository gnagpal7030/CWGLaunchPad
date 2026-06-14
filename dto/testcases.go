package dto

import "time"

type TestCase struct {
	ID           int       `json:"id"`
	QuestionID   int       `json:"question_id"`
	InputData    string    `json:"input_data"`
	ExpectedData string    `json:"expected_data"`
	IsHidden     bool      `json:"is_hidden"`
	CreatedAt    time.Time `json:"created_at"`
}
