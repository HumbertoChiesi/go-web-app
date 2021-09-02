package controllers

import "net/http"

//CreateUser inserts an user in the DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating an user!!"))
}

//SearchUsers searches for all the users in the DB
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all the users!!"))
}

//SearchUser searches for an specific user in the DB (by the user's ID)
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching an user!!"))
}

//UpdateUser update an user's content in the DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating an user!!"))
}

//DeleteUser deletes an user from the DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an user!!"))
}
