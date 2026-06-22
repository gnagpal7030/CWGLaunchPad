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

	fetchSubmission = `SELECT id, student_rollno, student_name, test_id, question_id, passed_testcases, total_testcases, submitted_at FROM submissions WHERE test_id = ? ORDER BY passed_testcases DESC`

	fetchSubmissionByRollNo = `SELECT
		id,
		student_rollno,
		student_name,
		test_id,
		question_id,
		passed_testcases,
		total_testcases,
		submitted_at
	FROM (
		SELECT
			s.*,
			ROW_NUMBER() OVER (
				PARTITION BY question_id
				ORDER BY passed_testcases DESC, submitted_at DESC
			) AS rn
		FROM submissions s
		WHERE s.test_id = ?
		AND s.student_rollno = ?
	) ranked
	WHERE rn = 1
	ORDER BY question_id;`
)

func (s *SubmissionRepo) InsertSubmission(submissionPayload *dto.SubmissionPayload) error {
	_, err := s.DB.Exec(insertIntoSubmission, &submissionPayload.StudentRollNo, &submissionPayload.StudentName, &submissionPayload.TestID, &submissionPayload.QuestionID, &submissionPayload.SourceCode, &submissionPayload.PassedTestcases, &submissionPayload.TotalTestcases, &submissionPayload.SubmittedAt)

	return err
}

func (s *SubmissionRepo) FetchSubmissions(testID int) ([]*dto.SubmissionPayload, error) {
	rows, err := s.DB.Query(fetchSubmission, testID)

	var submissions []*dto.SubmissionPayload

	for rows.Next() {
		submission := &dto.SubmissionPayload{}

		if err := rows.Scan(&submission.ID, &submission.StudentRollNo, &submission.StudentName, &submission.TestID, &submission.QuestionID, &submission.PassedTestcases, &submission.TotalTestcases, &submission.SubmittedAt); err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}

	return submissions, err
}

func (s *SubmissionRepo) FetchSubmissionsByRollNo(testID int, rollNo string) ([]*dto.SubmissionPayload, error) {
	rows, err := s.DB.Query(fetchSubmissionByRollNo, testID, rollNo)

	var submissions []*dto.SubmissionPayload

	for rows.Next() {
		submission := &dto.SubmissionPayload{}

		if err := rows.Scan(&submission.ID, &submission.StudentRollNo, &submission.StudentName, &submission.TestID, &submission.QuestionID, &submission.PassedTestcases, &submission.TotalTestcases, &submission.SubmittedAt); err != nil {
			return nil, err
		}
		submissions = append(submissions, submission)
	}

	return submissions, err
}
