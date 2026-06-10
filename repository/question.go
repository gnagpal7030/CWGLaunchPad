package repository

import (
	"CWDLaunchPad/model"
	"database/sql"
)

// sql queries to insert the questions in questions table

const (
	insertIntoQuestion = `
		INSERT INTO questions (
			title,
			description,
			constraints,
			starter_code,
			created_at,
			method_name,
			return_type,
			parameter_types,
			parameter_names,
			class_name
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
)

type QuestionRepository struct {
	DB *sql.DB
}

func (q *QuestionRepository) CreateQuestion(question *model.CreateQuestion) error {
	_, err := q.DB.Exec(
		insertIntoQuestion,
		question.Title,
		question.Description,
		question.Constraints,
		question.StarterCode,
		question.CreatedAt,
		question.MethodName,
		question.ReturnType,
		question.ParameterTypes,
		question.ParameterNames,
		question.ClassName,
	)

	return err
}
