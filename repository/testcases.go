package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/model"
	"database/sql"
)

const (
	insertTestCase = `INSERT INTO test_cases(question_id, input_data, expected_output, is_hidden, created_at) VALUES(?, ?, ?, ?, ?)`
)

type TestCaseRepository struct {
	DB *sql.DB
}

func GetTestCaseRepo() *TestCaseRepository {
	return &TestCaseRepository{
		DB: config.DB,
	}
}

func (t *TestCaseRepository) CreateTestCase(testCase *model.TestCase, questionID string) error {

	// TODO: can have a validation to check if the question_id exists or not

	_, err := t.DB.Exec(insertTestCase, questionID, testCase.InputData, testCase.ExpectedData, testCase.IsHidden, testCase.CreatedAt)
	return err
}
