package basic_api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{}
var mu sync.Mutex // manejar la concurrencia

func getUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func RunApi() {
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/create", createUser)

	fmt.Println("Server is running on http://localhost:8097")
	if err := http.ListenAndServe(":8097", nil); err != nil {
		fmt.Println(err)
	}
}
