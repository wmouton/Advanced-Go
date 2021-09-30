package main

import (
	"fmt"
	"github.com/Advanced-Go/Day-6/3-Domain-Driver-Design/easy-issues/application"
	"github.com/Advanced-Go/Day-6/3-Domain-Driver-Design/easy-issues/domain"
	"github.com/Advanced-Go/Day-6/3-Domain-Driver-Design/easy-issues/persistence/memory"
	"log"
	"net/http"
	"time"

	pb "github.com/Advanced-Go/Day-7/3-RPCs/easy-issues/protocol"

	"github.com/Advanced-Go/Day-7/3-RPCs/easy-issues/web"
	"google.golang.org/grpc"
	"net"
)

func main() {
	userRepo := memory.NewUserRepository()

	userService := application.UserService{
		UsersRepository: userRepo,
	}

	userController := web.UserController{
		UserService: userService,
	}

	for i := 0; i < 10; i += 1 {
		userService.Create(&domain.User{Name: fmt.Sprintf("User_%d", i)})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/", userController.List)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServer(grpcServer, userController)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go func() {
		grpcServer.Serve(lis)
	}()

	server := &http.Server{
		Addr:           ":8091",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
