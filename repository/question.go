package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/dto"
	"CWDLaunchPad/model"
	"database/sql"
	"fmt"
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
			param_types,
			param_names,
			class_name
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	fetchQuestions = `SELECT * FROM questions WHERE is_deleted=FALSE`
	deleteQuestion = `UPDATE questions SET is_deleted = TRUE WHERE id = ?`
	editQuestion   = `UPDATE questions SET title = ?, description = ?, constraints = ?, starter_code = ?, created_at = ?, method_name = ?, return_type = ?, param_types = ?, param_names = ?, class_name = ? WHERE id = ?`
)

type QuestionRepository struct {
	DB *sql.DB
}

func GetQuestionRepo() *QuestionRepository {
	return &QuestionRepository{
		DB: config.DB,
	}
}

func (q *QuestionRepository) CreateQuestion(question *model.Question) error {
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

func (q *QuestionRepository) GetQuestions(questionID ...string) ([]*model.Question, error) {

	var questions []*model.Question

	var query string = fetchQuestions
	if len(questionID) > 0 && questionID[0] != "" {
		query = fmt.Sprintf("%s AND id=%s", fetchQuestions, questionID[0])
	}

	res, err := q.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		q := &model.Question{}
		err := res.Scan(&q.ID, &q.Title, &q.Description, &q.Constraints, &q.StarterCode, &q.CreatedAt, &q.MethodName, &q.ReturnType, &q.ParameterTypes, &q.ParameterNames, &q.ClassName, &q.IsDeleted)
		if err != nil {
			fmt.Println("error fetching the question from db", err.Error())
			return nil, err
		}

		questions = append(questions, q)
	}
	return questions, nil
}

func (q *QuestionRepository) DeleteQuestion(questionID string) error {
	_, err := q.DB.Exec(deleteQuestion, questionID)
	return err
}

func (q *QuestionRepository) EditQuestion(question *dto.Question) error {
	_, err := q.DB.Exec(editQuestion, question.Title, question.Description, question.Constraints, question.StarterCode, question.CreatedAt, question.MethodName, question.ReturnType, question.ParameterTypes, question.ParameterNames, question.ClassName, question.ID)

	return err
}
