package controller

import (
	"encoding/json"
	"github.com/Advanced-Go/Day-6/3-Domain-Driver-Design/easy-issues/domain"
	"net/http"
)

// Controller for User model
type UserController struct {
	UserService domain.UserService
}

func (c UserController) List(w http.ResponseWriter, r *http.Request) {
	users, err := c.UserService.Users()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c UserController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c UserController) Delete(w http.ResponseWriter, r *http.Request) {
}
