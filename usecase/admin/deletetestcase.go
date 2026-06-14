package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteTestCaseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	testCaseID, ok := vars["testcase_id"]
	if !ok {
		fmt.Println("testcase_id parameter is required")
		http.Error(w, "testcase_id parameter is required", http.StatusBadRequest)
		return
	}

	// delete the testcase
	if err := AdminService.DeleteTestCase(testCaseID); err != nil {
		fmt.Println("error deleting testcase " + err.Error())
		http.Error(w, "error deleting testcase "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := &dto.Response{
		StatusCode: http.StatusOK,
		Message:    "resource delete successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		fmt.Println("error encoding response " + err.Error())
		http.Error(w, "error encoding response "+err.Error(), http.StatusInternalServerError)
		return
	}
}
