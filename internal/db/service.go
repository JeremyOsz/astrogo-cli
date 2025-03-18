package db

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// StartServer initializes and starts the database HTTP server
func StartServer(port int) error {
	// Initialize database
	if err := InitDB(); err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	defer CloseDB()

	// Start HTTP server
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/users", handleUsers)

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting database service on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleHealth checks if the service is running
func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database service is running"))
}

// handleUsers handles user-related HTTP requests
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetUsers(w, r)
	case "POST":
		handleCreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGetUsers handles GET /users requests
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// handleCreateUser handles POST /users requests
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" {
		http.Error(w, "Username and email are required", http.StatusBadRequest)
		return
	}

	newUser, err := CreateUser(user.Username, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
