package student

import (
	"CWDLaunchPad/dto"
	AdminService "CWDLaunchPad/service/adminservice"
	StudentService "CWDLaunchPad/service/studentservice"
	"encoding/json"
	"fmt"
	"net/http"
)

func TestJoinHandler(w http.ResponseWriter, r *http.Request) {

	// Accept Student Name and Roll No in the body
	var studentJoin *dto.StudentJoinRequestPayload
	if err := json.NewDecoder(r.Body).Decode(&studentJoin); err != nil {
		fmt.Println("error decoding the student join payload")
		http.Error(w, "error decoding the student join payload"+err.Error(), http.StatusBadRequest)
		return
	}

	// Store the student_id, student_name and testID in DB.
	if err := StudentService.StoreStudentJoinInfo(studentJoin.StudentName, studentJoin.StudentRollNo, studentJoin.TestID); err != nil {
		fmt.Println("error inserting the student data", err.Error())
		http.Error(w, "error inserting the student data"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the questions and every detail of it
	res, err := AdminService.GetSingleTest()
	if err != nil {
		fmt.Println("error getting the student join data", err.Error())
		http.Error(w, "error getting the student join data"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Code should be well formatted in the response
	response := &dto.StudentJoinResponse{
		Data:       res,
		Message:    "data fetched successfully",
		StatusCode: http.StatusOK,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		fmt.Println("error encoding the response", err.Error())
		http.Error(w, "error encoding the response", http.StatusInternalServerError)
	}
}
