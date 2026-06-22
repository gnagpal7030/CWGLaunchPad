package student

import (
	"CWDLaunchPad/dto"
	StudentService "CWDLaunchPad/service/studentservice"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testID, _ := strconv.Atoi(vars["test_id"])
	studentRollNo := vars["student_id"]
	res, err := StudentService.GetSubmissions(testID, studentRollNo)

	if err != nil {
		http.Error(w, "error fetching submissions"+err.Error(), http.StatusBadRequest)
		return
	}

	response := &dto.ResultsResponse{
		Data:       res,
		Message:    "data fetched sucessfully",
		StatusCode: http.StatusOK,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println("error enncoding the response" + err.Error())
		http.Error(w, "error enncoding the response"+err.Error(), http.StatusInternalServerError)
		return
	}
}
