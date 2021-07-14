package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/guilhermelopeseng/challenge-platform/application/repositories"
	"github.com/guilhermelopeseng/challenge-platform/application/usecases"
	"github.com/guilhermelopeseng/challenge-platform/framework/pb"
	"github.com/guilhermelopeseng/challenge-platform/framework/pb/servers"
	"github.com/guilhermelopeseng/challenge-platform/framework/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func main() {
	db = utils.ConnectDB()
	db.LogMode(true)

	port := flag.Int("port", 0, "Choose the server port")
	flag.Parse()
	log.Printf("Start server on port %v", *port)

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	adress := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Cannot start server: %v", err)
	}
}

func setUpUserServer() *servers.UserServer {
	userRepository := repositories.UserRepositoryDB{Db: db}
	userServer := servers.NewUserServer()
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userServer
}