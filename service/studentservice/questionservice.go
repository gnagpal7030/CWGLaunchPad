package StudentService

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/model"
	"CWDLaunchPad/repository"
	"fmt"
)

// All methods related to questions

func CreateQuestion(question *model.CreateQuestion) error {

	// add logic to create the question in the DB.

	questionRepo := &repository.QuestionRepository{
		DB: config.DB,
	}

	if err := questionRepo.CreateQuestion(question); err != nil {
		fmt.Println("error creating the question", err.Error())
		return err
	}

	return nil
}
