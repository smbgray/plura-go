package controllers

import "net/http"

// RegisterControllers expose /user/
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
