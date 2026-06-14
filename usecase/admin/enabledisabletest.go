package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func EnableDisableTest(w http.ResponseWriter, r *http.Request) {
	var payload *dto.EnableDisableTestPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {

		fmt.Println("error decoding the payload" + err.Error())
		http.Error(w, "error decoding the payload"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := AdminService.EnableDisableTest(payload); err != nil {
		fmt.Println("error updating the test" + err.Error())
		http.Error(w, "error updating the test"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&dto.Response{
		StatusCode: http.StatusOK,
		Message:    "resource udpated successfully",
	}); err != nil {
		fmt.Println("error enncoding the response" + err.Error())
		http.Error(w, "error enncoding the response"+err.Error(), http.StatusInternalServerError)
		return
	}

}
