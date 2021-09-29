package controller

import (
	"github.com/Advanced-Go/Day-6/3-Domain-Driver-Design/easy-issues/domain"
	"net/http"
)

// Controller for Project model
type ProjectController struct {
	ProjectService domain.ProjectService
}

func (c ProjectController) List(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c ProjectController) Delete(w http.ResponseWriter, r *http.Request) {

}
