package AdminService

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/repository"
)

func GetSubmissions(testID int) ([]*dto.SubmissionPayload, error) {
	return repository.GetSubmissionRepo().FetchSubmissions(testID)
}
