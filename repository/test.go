package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/dto"
	"database/sql"
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
