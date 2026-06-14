package AdminService

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/repository"
)

func CreateTest(test *dto.Test) error {
	return repository.GetTestRepo().CreateTest(test)
}

func EnableDisableTest(enableDisableTest *dto.EnableDisableTestPayload) error {
	return repository.GetTestRepo().EnableDisableTest(enableDisableTest)
}
