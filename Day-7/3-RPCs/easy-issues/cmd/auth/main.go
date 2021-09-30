package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/application"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/persistence/db"
	"github.com/Advanced-Go/Day-6/5-Authentication/easy-issues/web/handler"
	pb "github.com/Advanced-Go/Day-7/3-RPCs/easy-issues/protocol"
	"github.com/Advanced-Go/Day-7/3-RPCs/easy-issues/web"
	"google.golang.org/grpc"
)

const serverAddr = "127.0.0.1:10000"

func main() {
	authRepo := db.NewAuthRepository()
	authService := application.AuthService{
		AuthRepository: authRepo,
	}
	authController := web.AuthController{
		AuthService: authService,
		Secret:      application.Secret,
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	authController.UserClient = client

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/login", authController.Login)
	mux.HandleFunc("/auth/register", authController.Register)
	mux.HandleFunc("/auth/verify", handler.JWTAuthHandler(authController.Verify))

	server := &http.Server{
		Addr:           ":8090",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
