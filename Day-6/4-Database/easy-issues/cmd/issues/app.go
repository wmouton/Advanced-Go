package main

import (
	"fmt"
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/application"
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/domain"
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/persistence/db"
	"github.com/Advanced-Go/Day-6/4-Database/easy-issues/web/controller"
	"log"
	"net/http"
	"time"
)

func main() {
	issueRepo := db.NewIssueRepository()

	issueService := application.IssueService{
		IssueRepository: issueRepo,
	}

	issueController := controller.IssueController{
		IssueService: issueService,
	}

	for i := 0; i < 10; i += 1 {
		issueService.Create(
			&domain.Issue{
				Title:       fmt.Sprintf("Issue_%d", i),
				Description: "Task1",
				OwnerId:     1,
				ProjectId:   1,
				Status:      domain.StatusDone,
				Priority:    domain.PriorityHigh,
			})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", issueController.List)

	server := &http.Server{
		Addr:           ":8092",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
