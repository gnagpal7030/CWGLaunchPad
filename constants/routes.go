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

	// Student routes
	StudentJoin = "join"
)
