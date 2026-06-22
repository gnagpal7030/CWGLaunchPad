package studentservice

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/repository"
)

func GetSubmissions(testID int, rollNo string) ([]*dto.SubmissionPayload, error) {
	return repository.GetSubmissionRepo().FetchSubmissionsByRollNo(testID, rollNo)
}
