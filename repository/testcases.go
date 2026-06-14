package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/dto"
	"database/sql"
)

const (
	insertTestCase = `INSERT INTO test_cases(question_id, input_data, expected_output, is_hidden, created_at) VALUES(?, ?, ?, ?, ?)`
	deleteTestCase = `DELETE FROM test_cases WHERE id = ?`
	editTestCase   = `UPDATE test_cases SET input_data = ?, expected_output = ?, is_hidden = ?, created_at = ? WHERE id = ?`
	fetchtestCases = `SELECT * FROM test_cases`
)

type TestCaseRepository struct {
	DB *sql.DB
}

func GetTestCaseRepo() *TestCaseRepository {
	return &TestCaseRepository{
		DB: config.DB,
	}
}

func (t *TestCaseRepository) CreateTestCase(testCase *dto.TestCase, questionID string) error {

	// TODO: can have a validation to check if the question_id exists or not

	_, err := t.DB.Exec(insertTestCase, questionID, testCase.InputData, testCase.ExpectedData, testCase.IsHidden, testCase.CreatedAt)
	return err
}

func (t *TestCaseRepository) DeleteTestCase(testCaseID string) error {
	_, err := t.DB.Exec(deleteTestCase, testCaseID)
	return err
}

func (t *TestCaseRepository) EditTestCase(testCase *dto.TestCase) error {
	_, err := t.DB.Exec(editTestCase, testCase.InputData, testCase.ExpectedData, testCase.IsHidden, testCase.CreatedAt, testCase.ID)
	return err
}

func (q *QuestionRepository) getTestCases() ([]*dto.TestCase, error) {

	rows, err := q.DB.Query(fetchtestCases)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testCases []*dto.TestCase

	for rows.Next() {

		tc := &dto.TestCase{}

		err := rows.Scan(
			&tc.ID,
			&tc.QuestionID,
			&tc.InputData,
			&tc.ExpectedData,
			&tc.IsHidden,
			&tc.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		testCases = append(testCases, tc)
	}

	return testCases, nil
}
