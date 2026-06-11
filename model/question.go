package model

import "time"

type Question struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Constraints    string    `json:"constraints"`
	StarterCode    string    `json:"starter_code"`
	CreatedAt      time.Time `json:"created_at"`
	MethodName     string    `json:"method_name"`
	ReturnType     string    `json:"return_type"`
	ParameterTypes string    `json:"param_types"`
	ParameterNames string    `json:"param_names"`
	ClassName      string    `json:"class_name"`
	IsDeleted      bool      `json:"is_deleted"`
}
