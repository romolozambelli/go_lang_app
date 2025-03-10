package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// Load the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecTemplate(w, "login.html", nil)
}
