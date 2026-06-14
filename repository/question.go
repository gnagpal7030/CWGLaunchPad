package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/dto"
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

func (q *QuestionRepository) CreateQuestion(question *dto.Question) error {
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

func (q *QuestionRepository) getQuestions() ([]*dto.Question, error) {

	rows, err := q.DB.Query(fetchQuestions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*dto.Question

	for rows.Next() {

		question := &dto.Question{}

		err := rows.Scan(
			&question.ID,
			&question.Title,
			&question.Description,
			&question.Constraints,
			&question.StarterCode,
			&question.CreatedAt,
			&question.MethodName,
			&question.ReturnType,
			&question.ParameterTypes,
			&question.ParameterNames,
			&question.ClassName,
			&question.IsDeleted,
		)
		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func (q *QuestionRepository) GetQuestions(questionID ...string) ([]*dto.Question, error) {

	questions, err := q.getQuestions()
	if err != nil {
		return nil, err
	}

	testCases, err := q.getTestCases()
	if err != nil {
		return nil, err
	}

	// question_id -> []testcases
	testCaseMap := make(map[int][]*dto.TestCase)

	for _, tc := range testCases {
		testCaseMap[tc.QuestionID] = append(
			testCaseMap[tc.QuestionID],
			tc,
		)
	}

	for _, question := range questions {
		question.TestCases = testCaseMap[question.ID]
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
