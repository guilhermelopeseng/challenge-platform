package usecases

import (
	"log"

	"github.com/guilhermelopeseng/challenge-platform/application/repositories"
	"github.com/guilhermelopeseng/challenge-platform/domain"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *domain.User) (*domain.User, error) {
	user, err := u.UserRepository.Insert(user)
	if err != nil {
		log.Fatalf("Error to persist new user: %v", err)
		return user, err
	}
	return user, nil
}