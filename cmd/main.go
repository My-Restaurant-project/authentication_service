package main

import (
	config "authentication_service/configs"
	pb "authentication_service/genproto/authentication_service"
	"authentication_service/pkg"
	"authentication_service/repositories"
	"authentication_service/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	config := config.Load()

	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()
	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer listener.Close()

	log.Printf("Server started on port " + config.URL_PORT)

	authStorage := repositories.NewUserRepository(db)

	as := services.NewAuthService(authStorage)

	s := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(s, as)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
