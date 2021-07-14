package repositories

import (
	"github.com/guilhermelopeseng/challenge-platform/domain"
	"github.com/jinzhu/gorm"
	"log"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDB struct {
	Db *gorm.DB
}

func (repo UserRepositoryDB) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()
	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return user, err
	}

	err = repo.Db.Create(user).Error
	if err != nil {
		log.Fatalf("Error to persis user: %v", err)
		return user, err
	}
	
	return user, nil
}
