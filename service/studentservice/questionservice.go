package AdminService

import (
	"CWDLaunchPad/model"
	"CWDLaunchPad/repository"
	"fmt"
)

// All methods related to questions

func CreateQuestion(question *model.Question) error {

	// add logic to create the question in the DB.
	questionRepo := repository.GetQuestionRepo()
	if err := questionRepo.CreateQuestion(question); err != nil {
		fmt.Println("error creating the question", err.Error())
		return err
	}

	return nil
}

func GetQuestions(questionID ...string) ([]*model.Question, error) {
	questionRepo := repository.GetQuestionRepo()
	return questionRepo.GetQuestions(questionID...)
}

func DeleteQuestion(questionID string) error {
	questionRepo := repository.GetQuestionRepo()
	return questionRepo.DeleteQuestion(questionID)
}
