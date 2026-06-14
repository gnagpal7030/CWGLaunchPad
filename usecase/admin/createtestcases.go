package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateTestCases(w http.ResponseWriter, r *http.Request) {
	// accept question_id for mapping from path parameters
	vars := mux.Vars(r)
	questionID, ok := vars["id"]
	if !ok {
		fmt.Println("id parameter is required")
		http.Error(w, "id parameter. is required", http.StatusBadRequest)
		return
	}

	// create testcase for question_id
	var testCase *dto.TestCase
	if err := json.NewDecoder(r.Body).Decode(&testCase); err != nil {
		fmt.Println("error decoding the test case payload" + err.Error())
		http.Error(w, "error decoding the test case payload", http.StatusBadRequest)
		return
	}

	if err := AdminService.CreateTestCase(testCase, questionID); err != nil {
		fmt.Println("error creating test case" + err.Error())
		http.Error(w, "error creating the test case", http.StatusInternalServerError)
		return
	}

	response := &dto.Response{
		StatusCode: http.StatusCreated,
		Message:    "resource created successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
