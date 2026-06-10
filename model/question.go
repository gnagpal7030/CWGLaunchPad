package model

import "time"

type CreateQuestion struct {
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
}
