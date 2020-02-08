package controllers

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list users"))
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An user"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update user"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}