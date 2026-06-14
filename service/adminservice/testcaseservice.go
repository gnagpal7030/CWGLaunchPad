package AdminService

import (
	"CWDLaunchPad/model"
	"CWDLaunchPad/repository"
)

func CreateTestCase(testCase *model.TestCase, questionID string) error {
	return repository.GetTestCaseRepo().CreateTestCase(testCase, questionID)
}
