package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTestHandler(w http.ResponseWriter, r *http.Request) {
	res, err := AdminService.GetAllTests()
	if err != nil {
		fmt.Println("error getting data", err.Error())
		if err.Error() == "no data found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "error getting data"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response := &dto.GetTestResponse{
		Data:       res,
		Message:    "data fetched successfully",
		StatusCode: http.StatusOK,
	}

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "error encoding the result", http.StatusInternalServerError)
	}
}
