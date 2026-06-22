package constants

import "fmt"

func GetRoute(routeName string) string {
	return fmt.Sprintf("/api/v1/%s", routeName)
}

const (

	// Admin routes
	AdminLogin = "login"
	Question   = "question"
	TestCases  = "testcases"
	Tests      = "tests"
	Results    = "results"

	// Student routes
	StudentJoin = "join"
	Submit      = "submit"
)
