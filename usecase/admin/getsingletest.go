package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetSingleTestHandler(w http.ResponseWriter, r *http.Request) {

	res, err := AdminService.GetSingleTest()
	if err != nil {
		fmt.Println("error fetching the test details", err.Error())
		http.Error(w, "error fetching the test details"+err.Error(), http.StatusInternalServerError)
		return
	}

	response := &dto.SingleTestResponse{
		Data:       res,
		Message:    "data fetched successfully",
		StatusCode: http.StatusOK,
	}

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "error encoding the result", http.StatusInternalServerError)
	}
}
