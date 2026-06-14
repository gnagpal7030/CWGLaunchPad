package usecase

import (
	"CWDLaunchPad/dto"
	"CWDLaunchPad/model"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {

	var question *dto.Question

	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		fmt.Println("error decoding the request", err.Error())
		http.Error(w, "error decoding the request body", http.StatusBadRequest)
		return
	}

	// TODO: Can add validations in future if required

	if err := AdminService.CreateQuestion((*model.Question)(question)); err != nil {
		fmt.Println("error creating the question", err.Error())
		http.Error(w, "error creating the question", http.StatusInternalServerError)
		return
	}

	// send success response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	response := &dto.Response{
		StatusCode: http.StatusCreated,
		Message:    "Question created successfully",
	}

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		fmt.Println("error encoding the response", err.Error())
		http.Error(w, "error encoding the response", http.StatusInternalServerError)
	}
}
