package usecase

import (
	"CWDLaunchPad/constants"
	"CWDLaunchPad/dto"
	"CWDLaunchPad/middleware"
	"encoding/json"
	"net/http"
	"os"
)

func AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	// get username and password from request
	var req *dto.LoginPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// verify the username and password
	if req.UserName == "" || req.Password == "" {
		http.Error(w, "UserName and Password is required", http.StatusBadRequest)
		return
	}

	adminUsername := os.Getenv(constants.AdminUserName)
	adminPassword := os.Getenv(constants.AdminPassword)

	if req.UserName != adminUsername || req.Password != adminPassword {
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

	// send back the details with jwt token
	tokenString, err := middleware.CreateToken(adminUsername)
	if err != nil {
		http.Error(w, "error creating the token for admin"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := &dto.LoginResponse{
		Token: tokenString,
	}

	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "error sending response", http.StatusInternalServerError)
		return
	}

}
