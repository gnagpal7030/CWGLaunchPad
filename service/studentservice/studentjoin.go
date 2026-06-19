package studentservice

import "CWDLaunchPad/repository"

func StoreStudentJoinInfo(studentName, studentRollNo string, testID int) error {
	return repository.GetTestRepo().InsertStudentJoinData(studentRollNo, studentName, testID)
}
