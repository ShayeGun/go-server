package service

import (
	"github.com/ShayeGun/go-server/internal/common"
	"github.com/ShayeGun/go-server/models"
)

type UserService struct {
	UserRepository common.UserRepositoryInterface
}

func NewUserService(ur common.UserRepositoryInterface) common.UserServiceInterface {
	us := &UserService{
		UserRepository: ur,
	}

	return us
}

// NOTE: currently service looks like wrapper around repository its because business logic is not complicated in future it might change

func (us *UserService) GetUser(uid string) (models.User, error) {
	return us.UserRepository.GetById(uid)
}

func (us *UserService) AddUser(u models.User) (models.User, error) {
	return us.UserRepository.Add(u)
}

func (us *UserService) UpdateUser(u models.User) (models.User, error) {
	return us.UserRepository.Update(u)
}

func (us *UserService) DeleteUser(uid string) error {
	return us.UserRepository.Delete(uid)
}
