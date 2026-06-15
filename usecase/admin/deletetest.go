package usecase

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteTestHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testID := params["test_id"]

	if err := AdminService.DeleteTest(testID); err != nil {
		fmt.Println("error deleting the test" + err.Error())
		http.Error(w, "error deleting the test"+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&dto.Response{
		StatusCode: http.StatusOK,
		Message:    "resource deleted successfully",
	}); err != nil {
		fmt.Println("error encoding the response" + err.Error())
		http.Error(w, "error encoding the response"+err.Error(), http.StatusInternalServerError)
		return
	}

}
