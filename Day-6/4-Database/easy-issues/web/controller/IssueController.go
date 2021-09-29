package controller

import (
	"encoding/json"
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/domain"
	"net/http"
)

// Controller for Issue model
type IssueController struct {
	IssueService domain.IssueService
}

func (c IssueController) List(w http.ResponseWriter, r *http.Request) {
	issues, err := c.IssueService.Issues()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	issuesJson, err := json.Marshal(issues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(issuesJson)
}

func (c IssueController) Show(w http.ResponseWriter, r *http.Request) {

}

func (c IssueController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c IssueController) Delete(w http.ResponseWriter, r *http.Request) {

}
