package controllers

import "net/http"

//CreateUser creates a new user in the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User!"))
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