package dto

type StudentJoinRequestPayload struct {
	StudentName   string `json:"student_name"`
	StudentRollNo string `json:"student_rollno"`
	TestID        int    `json:"test_id"`
}

// type StudentJoinData struct {
// 	QuestionData  Question `json:"question_data"`
// 	TestCasesData TestCase `json:"testcases_data"`
// 	TestData      Test     `json:"test_data"`
// }

type StudentJoinResponse struct {
	Data       *SingleTest `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}
