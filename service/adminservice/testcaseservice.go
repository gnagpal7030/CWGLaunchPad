package AdminService

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/repository"
)

func CreateTestCase(testCase *dto.TestCase, questionID string) error {
	return repository.GetTestCaseRepo().CreateTestCase(testCase, questionID)
}

func DeleteTestCase(testCaseID string) error {
	return repository.GetTestCaseRepo().DeleteTestCase(testCaseID)
}

func EditTestCase(testCase *dto.TestCase) error {
	return repository.GetTestCaseRepo().EditTestCase(testCase)
}
