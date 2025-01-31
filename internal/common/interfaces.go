package common

import (
	"github.com/ShayeGun/go-server/models"
)

type UserRepositoryInterface interface {
	Add(models.User) (models.User, error)
	GetById(string) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(string) error
}

type UserServiceInterface interface {
	GetUser(uid string) (models.User, error)
	AddUser(u models.User) (models.User, error)
	UpdateUser(u models.User) (models.User, error)
	DeleteUser(uid string) error
}

type RepositoryInterface interface {
	GetUserTable() UserRepositoryInterface
}

type ServiceInterface interface {
	GetUserService() UserServiceInterface
}
