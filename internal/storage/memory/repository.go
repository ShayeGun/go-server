package memory

import (
	"github.com/ShayeGun/go-server/internal/service"
	"github.com/ShayeGun/go-server/models"
)

type Repository struct {
	userCollection service.UserRepositoryInterface
}

func NewRepository() *Repository {
	return &Repository{
		userCollection: &UserRepository{
			users: make(map[string]models.User),
		},
	}
}

func (r *Repository) GetUserTable() service.UserRepositoryInterface {
	return r.userCollection
}
