package repository

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/dto"
	"database/sql"
)

type SubmissionRepo struct {
	DB *sql.DB
}

func GetSubmissionRepo() *SubmissionRepo {
	return &SubmissionRepo{
		DB: config.DB,
	}
}

const (
	insertIntoSubmission = `INSERT INTO submissions(student_rollno, student_name, test_id, question_id, source_code, passed_testcases, total_testcases, submitted_at) VALUES(?,?,?,?,?,?,?,?)`
)

func (s SubmissionRepo) InsertSubmission(submissionPayload *dto.SubmissionPayload) error {
	_, err := s.DB.Exec(insertIntoSubmission, &submissionPayload.StudentRollNo, &submissionPayload.StudentName, &submissionPayload.TestID, &submissionPayload.QuestionID, &submissionPayload.SourceCode, &submissionPayload.PassedTestcases, &submissionPayload.TotalTestcases, &submissionPayload.SubmittedAt)

	return err
}
