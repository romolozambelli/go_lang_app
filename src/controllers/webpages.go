package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// Load the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}

// Load the register user page
func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "register_user.html", nil)
}
