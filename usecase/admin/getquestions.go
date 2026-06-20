package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetQuestionsHandler(w http.ResponseWriter, r *http.Request) {

	questionID := r.URL.Query().Get("id")
	questionIDInt, _ := strconv.Atoi(questionID)
	res, err := AdminService.GetQuestions(questionIDInt)
	if err != nil {
		http.Error(w, "error fetching questions from DB"+err.Error(), http.StatusInternalServerError)
		return
	}

	// send all questions
	var questionResponse []*dto.Question
	for _, q := range res {
		questionResponse = append(questionResponse, q)
	}

	message := "data fetched successfully"
	statusCode := http.StatusOK
	if len(questionResponse) == 0 {
		message = "no data found"
		statusCode = http.StatusNotFound
	}

	response := &dto.GetQuestionsResponse{
		Data:       questionResponse,
		StatusCode: statusCode,
		Message:    message,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "error encoding the result", http.StatusInternalServerError)
	}
}
