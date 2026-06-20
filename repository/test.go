package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/dto"
	"database/sql"
	"errors"
)

type TestRepository struct {
	DB *sql.DB
}

func GetTestRepo() *TestRepository {
	return &TestRepository{
		DB: config.DB,
	}
}

const (
	insertIntoTest                = `INSERT INTO tests(title, duration_minutes, description, enabled, created_at) VALUES(?, ?, ?, ?, ?)`
	updateEnabledStatus           = `UPDATE tests SET enabled = ? WHERE id = ?`
	insertIntoTestQuestionMapping = `INSERT INTO test_questions(test_id, question_id) VALUES (?, ?)`
	fetchTests                    = `SELECT * FROM tests`
	fetchTestWithQuestions        = `SELECT q.* FROM questions q JOIN test_questions tq ON tq.question_id = q.id WHERE tq.test_id = ? AND q.is_deleted = FALSE`
	deleteTest                    = `DELETE FROM tests WHERE id = ?`
	insertIntoStudentTestMapping  = `INSERT into student_test_mapping (student_name, roll_no, test_id) VALUES(?, ?, ?)`
)

func (t *TestRepository) CreateTest(test *dto.Test) error {
	res, err := t.DB.Exec(insertIntoTest, test.Title, test.DurationMinutes, test.Description, test.Enabled, test.CreatedAt)

	// Also add the mapping between test ID and questions ID
	testID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	questionsIDs := test.QuestionIDs

	// Use testID for inserting mappings
	for _, questionID := range questionsIDs {
		_, err := t.DB.Exec(
			insertIntoTestQuestionMapping,
			testID,
			questionID,
		)
		if err != nil {
			return err
		}
	}

	return err
}

func (t *TestRepository) EnableDisableTest(payload *dto.EnableDisableTestPayload) error {
	_, err := t.DB.Exec(updateEnabledStatus, payload.Enable, payload.TestID)
	return err
}

func (t *TestRepository) GetAllTests() ([]*dto.Test, error) {
	rows, err := t.DB.Query(fetchTests)
	if err != nil {
		return nil, err
	}

	var tests []*dto.Test

	for rows.Next() {
		test := &dto.Test{}
		err := rows.Scan(&test.ID, &test.Title, &test.DurationMinutes, &test.Description, &test.Enabled, &test.CreatedAt)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}

	if len(tests) == 0 {
		return nil, errors.New("no data found")
	}

	return tests, err
}

func (t *TestRepository) GetSingleTest(testID int) (*dto.SingleTest, error) {

	// Fetch test details
	test := &dto.Test{}

	err := t.DB.QueryRow(
		fetchTests+` WHERE ID = ?`,
		testID,
	).Scan(
		&test.ID,
		&test.Title,
		&test.DurationMinutes,
		&test.Description,
		&test.Enabled,
		&test.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Fetch all questions associated with the test
	rows, err := t.DB.Query(
		fetchTestWithQuestions,
		testID,
	)

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

	return &dto.SingleTest{
		TestData:  test,
		Questions: questions,
	}, nil
}

func (t *TestRepository) DeleteTest(testID string) error {
	_, err := t.DB.Exec(deleteTest, testID)
	return err
}

func (t *TestRepository) InsertStudentJoinData(studentName, studentRollNo string, testID int) error {

	_, err := t.DB.Exec(insertIntoStudentTestMapping, studentName, studentRollNo, testID)

	return err
}
