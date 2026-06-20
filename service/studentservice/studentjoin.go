package studentservice

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/repository"
)

func StoreStudentJoinInfo(studentName, studentRollNo string, testID int) error {
	return repository.GetTestRepo().InsertStudentJoinData(studentName, studentRollNo, testID)
}

func GetTestCasesForQuestion(questionID int) ([]*dto.TestCase, error) {
	return repository.GetTestCaseRepo().GetTestCases(questionID)
}

func InsertSubmission(payload *dto.SubmissionPayload) error {
	return repository.GetSubmissionRepo().InsertSubmission(payload)
}
