package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//CreateUser creates a new user in the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}


	repository := repositories.NewUsersRepositories(db)
	UserID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID entered: %d", UserID)))
}

//SearchUsers Searches for all users saved in the database.
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for all users!"))
}

//SearchUser Searches for a user saved in the database.
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for a user!"))
}

//UpdateUser Changes the user information in the database.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}

//DeleteUser Deletes the user information in the database.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User!"))
}