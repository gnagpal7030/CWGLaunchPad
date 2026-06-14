package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func EditQuestionHandler(w http.ResponseWriter, r *http.Request) {

	// accept the payload - it will contain the id of the question

	var question *dto.Question

	// decode the payload
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, "error decoding the body"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := AdminService.EditQuestion(question); err != nil {
		http.Error(w, "error editing the question", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&dto.Response{
		StatusCode: http.StatusOK,
		Message:    "resource edited successfully",
	}); err != nil {
		fmt.Println("error encoding the result", err.Error())
		http.Error(w, "error decoding the result", http.StatusInternalServerError)
	}
}
