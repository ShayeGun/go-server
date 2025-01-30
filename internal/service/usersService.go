package service

import (
	"github.com/ShayeGun/go-server/models"
)

type userConfig func(us *UserService) error

type UserRepositoryInterface interface {
	Add(models.User) (models.User, error)
	GetById(string) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(string) error
}

type UserService struct {
	UserRepository UserRepositoryInterface
}

func NewUserService(cfgs ...userConfig) (*UserService, error) {
	us := &UserService{}

	for _, cfg := range cfgs {
		err := cfg(us)
		if err != nil {
			return nil, err
		}
	}

	return us, nil
}

func WithUserRepository(ur UserRepositoryInterface) userConfig {
	return func(us *UserService) error {
		us.UserRepository = ur
		return nil
	}
}
