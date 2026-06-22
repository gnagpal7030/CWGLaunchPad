package dto

import (
	"time"
)

type SubmissionPayload struct {
	ID              int       `json:"id"`
	StudentRollNo   string    `json:"student_rollno"`
	StudentName     string    `json:"student_name"`
	TestID          int       `json:"test_id"`
	QuestionID      int       `json:"question_id"`
	SourceCode      string    `json:"source_code"`
	PassedTestcases int       `json:"passed_testcases"`
	TotalTestcases  int       `json:"total_testcases"`
	SubmittedAt     time.Time `json:"submitted_at"`
}

type SubmitCodeResult struct {
	PassedTestcases int    `json:"passed_testcases"`
	TotalTestcases  int    `json:"total_testcases"`
	Message         string `json:"message"`
	StatusCode      int    `json:"status_code"`
	Error           string `json:"error,omitempty"`
}

type ResultsResponse struct {
	Data       []*SubmissionPayload `json:"data"`
	Message    string               `json:"message"`
	StatusCode int                  `json:"status_code"`
}
