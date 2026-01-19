package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// USER MODEL
type User struct {
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// In-memory storage (for demo)
var users = make(map[string]string)

// CORS  
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
}

//  SIGNUP API  
func signup(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Password == "" || (user.Email == "" && user.Mobile == "") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "email or mobile and password required",
		})
		return
	}

	key := user.Email
	if key == "" {
		key = user.Mobile
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	users[key] = string(hash)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "signup successful",
	})
}

//  LOGIN API 
func login(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	key := user.Email
	if key == "" {
		key = user.Mobile
	}

	hash, ok := users[key]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "user not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "invalid password",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "login successful",
	})
}

//  MAIN  
func main() {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
