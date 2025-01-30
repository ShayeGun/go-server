package memory

import (
	"errors"
	"sync"

	"github.com/ShayeGun/go-server/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type UserRepository struct {
	users map[string]models.User
	sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]models.User),
	}
}

func (r *UserRepository) GetById(id string) (models.User, error) {
	u, ok := r.users[id]
	if !ok {
		return models.User{}, ErrUserNotFound
	}

	return u, nil
}

func (r *UserRepository) Add(u models.User) (models.User, error) {
	if r.users == nil {
		r.Lock()
		r.users = make(map[string]models.User)
		r.Unlock()
	}

	if _, ok := r.users[u.GetID()]; ok {
		return models.User{}, ErrUserAlreadyExists
	}

	r.users[u.GetID()] = u
	return r.users[u.GetID()], nil
}

func (r *UserRepository) Update(u models.User) (models.User, error) {
	if r.users == nil {
		r.Lock()
		r.users = make(map[string]models.User)
		r.Unlock()
	}

	existedUser, ok := r.users[u.GetID()]

	if !ok {
		return models.User{}, ErrUserNotFound
	}

	existedUser.Password = u.GetPassword()
	r.users[u.GetID()] = existedUser

	return r.users[u.GetID()], nil
}

func (r *UserRepository) Delete(id string) error {
	if r.users == nil {
		r.Lock()
		r.users = make(map[string]models.User)
		r.Unlock()
	}

	if _, ok := r.users[id]; !ok {
		return ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}
