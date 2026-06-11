package main

import (
	"CWDLaunchPad/config"
	"CWDLaunchPad/constants"
	usecase "CWDLaunchPad/usecase/admin"
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

	// TODO: Add zap logger through the application

	// ----------------  Admin Routes ---------------
	adminRouter.HandleFunc(constants.GetRoute(constants.AdminLogin), usecase.AdminLoginHandler).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.GetRoute(constants.StudentQuestion), usecase.CreateQuestionHandler).Methods(http.MethodPost)
	adminRouter.HandleFunc(constants.GetRoute(constants.StudentQuestion), usecase.GetQuestionsHandler).Methods(http.MethodGet)
	adminRouter.HandleFunc(constants.GetRoute(constants.StudentQuestion), usecase.DeleteQuestionHandler).Methods(http.MethodDelete)
	adminRouter.HandleFunc(constants.GetRoute(constants.StudentQuestion), usecase.EditQuestionHandler).Methods(http.MethodPut)

	// ---------------- Student Routes --------------
	studentRouter.HandleFunc(constants.GetRoute(constants.StudentJoin), usecase.AdminLoginHandler).Methods(http.MethodPost)

	fmt.Println("Server is starting")

	port := os.Getenv(constants.AppPort)
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		fmt.Println("Error running the server", err.Error())
	}
}
