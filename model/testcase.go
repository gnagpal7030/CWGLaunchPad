package model

import "time"

type TestCase struct {
	ID           string    `json:"id"`
	QuestionID   string    `json:"question_id"`
	InputData    string    `json:"input_data"`
	ExpectedData string    `json:"expected_data"`
	IsHidden     bool      `json:"is_hidden"`
	CreatedAt    time.Time `json:"created_at"`
}
