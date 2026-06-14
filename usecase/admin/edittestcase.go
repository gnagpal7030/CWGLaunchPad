package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func EditTestCaseHandler(w http.ResponseWriter, r *http.Request) {
	var testCase *dto.TestCase
	if err := json.NewDecoder(r.Body).Decode(&testCase); err != nil {
		fmt.Println("error decoding the test case payload" + err.Error())
		http.Error(w, "error decoding the test case payload", http.StatusBadRequest)
		return
	}

	if err := AdminService.EditTestCase(testCase); err != nil {
		fmt.Println("error editing test case" + err.Error())
		http.Error(w, "error editing the test case", http.StatusInternalServerError)
		return
	}

	response := &dto.Response{
		StatusCode: http.StatusCreated,
		Message:    "resource edited successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}
