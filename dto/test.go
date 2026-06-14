package dto

import "time"

type Test struct {
	Title           string    `json:"title"`
	DurationMinutes int       `json:"duration_minutes"`
	Description     string    `json:"description"`
	Enabled         bool      `json:"enabled"`
	CreatedAt       time.Time `json:"created_at"`
	QuestionIDs     []int     `json:"question_ids"`
}

type EnableDisableTestPayload struct {
	TestID int  `json:"test_id"`
	Enable bool `json:"enable"`
}
