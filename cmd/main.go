package main

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/constants"
	"CWDLaunchPad/middleware"
	admin "CWDLaunchPad/usecase/admin"
	student "CWDLaunchPad/usecase/students"

	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// load env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("error loading .env file")
		return
	}

	// initliase mysql connection
	if err := config.InitliaseMySQLConnection(); err != nil {
		fmt.Println("error connecting to mysql db", err.Error())
		return
	}
	defer config.DB.Close()

	r := mux.NewRouter()

	adminRouter := r.PathPrefix("/admin").Subrouter()
	studentRouter := r.PathPrefix("/student").Subrouter()

	adminRouter.Use(middleware.ValidateAdminToken)
	// TODO: Add zap logger through the application

	// ----------------  Admin Routes Start ---------------
	r.HandleFunc(constants.GetRoute(constants.AdminLogin), admin.AdminLoginHandler).Methods(http.MethodPost)

	// ---------------- Questions Routes -------------
	adminRouter.HandleFunc(constants.GetRoute(constants.Question), admin.CreateQuestionHandler).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.GetRoute(constants.Question), admin.GetQuestionsHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc(constants.GetRoute(constants.Question), admin.DeleteQuestionHandler).Methods(http.MethodDelete)
	adminRouter.HandleFunc(constants.GetRoute(constants.Question), admin.EditQuestionHandler).Methods(http.MethodPut)

	// ---------------- Test Cases Routes ------------
	adminRouter.HandleFunc(constants.GetRoute(constants.TestCases)+"/{id}", admin.CreateTestCases).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.GetRoute(constants.TestCases)+"/{testcase_id}", admin.DeleteTestCaseHandler).Methods(http.MethodDelete)
	adminRouter.HandleFunc(constants.GetRoute(constants.TestCases)+"/{testcase_id}", admin.EditTestCaseHandler).Methods(http.MethodPut)

	// ----------------- Tests Routes ----------------
	adminRouter.HandleFunc((constants.GetRoute(constants.Tests)), admin.CreateTestHandler).Methods(http.MethodPost)
	adminRouter.HandleFunc((constants.GetRoute(constants.Tests)), admin.EnableDisableTest).Methods(http.MethodPut)
	adminRouter.HandleFunc((constants.GetRoute(constants.Tests)), admin.GetTestHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc((constants.GetRoute(constants.Tests))+"/{test_id}", admin.GetSingleTestHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc(constants.GetRoute(constants.Tests)+"/{test_id}", admin.DeleteTestHandler).Methods(http.MethodDelete)

	// ---------------- Admin Routes End ------------

	// ---------------- Student Routes Start--------------
	studentRouter.HandleFunc(constants.GetRoute(constants.StudentJoin), student.TestJoinHandler).Methods(http.MethodPost)
	studentRouter.HandleFunc(constants.GetRoute(constants.Submit), student.CodeSubmissionHandler).Methods(http.MethodPost)

	// ---------------- Student Routes End----------------

	fmt.Println("Server is starting")

	port := os.Getenv(constants.AppPort)
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		fmt.Println("Error running the server", err.Error())
	}
}
