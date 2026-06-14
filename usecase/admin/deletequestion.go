package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	questionID := r.URL.Query().Get("id")

	if questionID == "" {
		// without questionID, delete operation is not possible.
		http.Error(w, "questionID is mandatory", http.StatusBadRequest)
		return
	}

	if err := AdminService.DeleteQuestion(questionID); err != nil {
		http.Error(w, "error deleting the question", http.StatusInternalServerError)
		return
	}

	response := &dto.Response{
		StatusCode: http.StatusOK,
		Message:    "resource delete successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		fmt.Println("error encoding the response")
		http.Error(w, "error encoding the response", http.StatusInternalServerError)
	}
}
