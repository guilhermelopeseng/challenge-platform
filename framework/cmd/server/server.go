package main

import (
	"fmt"
	"log"

	"github.com/guilhermelopeseng/challenge-platform/application/repositories"
	"github.com/guilhermelopeseng/challenge-platform/domain"
	"github.com/guilhermelopeseng/challenge-platform/framework/utils"
	_ "github.com/lib/pq"
)

func main() {
	db := utils.ConnectDB()

	user := domain.User{
		Name: "Guilherme",
		Email: "gui@test.com",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDB{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Name, result.Email, result.Token)
}