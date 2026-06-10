package constants

import "fmt"

func GetRoute(routeName string) string {
	return fmt.Sprintf("/api/v1/%s", routeName)
}

const (

	// Admin routes
	AdminLogin = "login"

	// Student routes
	StudentJoin           = "join"
	StudentCreateQuestion = "question"
)
