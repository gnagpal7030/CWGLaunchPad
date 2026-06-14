package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTestHandler(w http.ResponseWriter, r *http.Request) {

	var test *dto.Test
	if err := json.NewDecoder(r.Body).Decode(&test); err != nil {
		fmt.Println("error decoding the payload" + err.Error())
		http.Error(w, "error decoding the payload"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := AdminService.CreateTest(test); err != nil {
		fmt.Println("error creating the test" + err.Error())
		http.Error(w, "error creating the test"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&dto.Response{
		StatusCode: http.StatusCreated,
		Message:    "resource created successfully",
	}); err != nil {
		fmt.Println("error enncoding the response" + err.Error())
		http.Error(w, "error enncoding the response"+err.Error(), http.StatusInternalServerError)
		return
	}
}
