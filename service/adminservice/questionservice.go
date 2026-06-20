package AdminService

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/repository"
	"fmt"
)

// All methods related to questions

func CreateQuestion(question *dto.Question) error {

	// add logic to create the question in the DB.
	questionRepo := repository.GetQuestionRepo()
	if err := questionRepo.CreateQuestion(question); err != nil {
		fmt.Println("error creating the question", err.Error())
		return err
	}

	return nil
}

func GetQuestions(questionID ...int) ([]*dto.Question, error) {
	return repository.GetQuestionRepo().GetQuestions(questionID...)
}

func DeleteQuestion(questionID string) error {
	return repository.GetQuestionRepo().DeleteQuestion(questionID)
}

func EditQuestion(question *dto.Question) error {
	return repository.GetQuestionRepo().EditQuestion(question)
}
