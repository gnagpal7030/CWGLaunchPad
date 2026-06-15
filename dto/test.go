package dto

import "time"

type Test struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	DurationMinutes int       `json:"duration_minutes"`
	Description     string    `json:"description"`
	Enabled         bool      `json:"enabled"`
	CreatedAt       time.Time `json:"created_at"`
	QuestionIDs     []int     `json:"question_ids,omitempty"`
}

type EnableDisableTestPayload struct {
	TestID int  `json:"test_id"`
	Enable bool `json:"enable"`
}

type GetTestResponse struct {
	Data       []*Test `json:"data"`
	Message    string  `json:"message"`
	StatusCode int     `json:"status_code"`
}

type SingleTest struct {
	TestData  *Test       `json:"test_data"`
	Questions []*Question `json:"questions"`
}

type SingleTestResponse struct {
	Data       *SingleTest `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}
