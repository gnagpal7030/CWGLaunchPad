package dto

import "time"

// Questions struct

type Question struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Constraints    string    `json:"constraints"`
	StarterCode    string    `json:"starter_code"`
	CreatedAt      time.Time `json:"created_at"`
	MethodName     string    `json:"method_name"`
	ReturnType     string    `json:"return_type"`
	ParameterTypes string    `json:"parameter_types"`
	ParameterNames string    `json:"parameter_names"`
	ClassName      string    `json:"class_name"`
	IsDeleted      bool      `json:"is_deleted"`
}

type CreateQuestionResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GetQuestionsResponse struct {
	Data       []*Question `json:"data,omitempty"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
}
