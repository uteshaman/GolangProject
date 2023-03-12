package main

import (
	"fmt"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func main() {
    router := mux.NewRouter()

    // Auth endpoints
    router.HandleFunc("/register", registerUser).Methods("POST")
    router.HandleFunc("/login", login).Methods("POST")
    router.HandleFunc("/logout", logout).Methods("POST")

    // Todo list endpoints
    router.HandleFunc("/todos", getTodoLists).Methods("GET")
    router.HandleFunc("/todos/{id}", getTodoList).Methods("GET")
    router.HandleFunc("/todos", createTodoList).Methods("POST")
    router.HandleFunc("/todos/{id}", updateTodoList).Methods("PUT")
    router.HandleFunc("/todos/{id}", deleteTodoList).Methods("DELETE")

    // Todo endpoints
    router.HandleFunc("/todos/{listId}/items", getTodos).Methods("GET")
    router.HandleFunc("/todos/{listId}/items/{id}", getTodo).Methods("GET")
    router.HandleFunc("/todos/{listId}/items", createTodo).Methods("POST")
    router.HandleFunc("/todos/{listId}/items/{id}", updateTodo).Methods("PUT")
    router.HandleFunc("/todos/{listId}/items/{id}", deleteTodo).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}
func registerUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // Check if user already exists
    // ...

    // Hash user's password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    // Store user in database
    // ...

    // Return success message to client
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
func login(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    // Retrieve user from database
    // ...

    // Compare user's password to hashed password in database
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(storedPassword))
    if err != nil {
        http.Error
