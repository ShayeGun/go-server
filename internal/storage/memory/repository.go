package memory

import (
	"github.com/ShayeGun/go-server/internal/common"
	"github.com/ShayeGun/go-server/models"
)

type Repository struct {
	userCollection common.UserRepositoryInterface
}

func NewRepository() *Repository {
	return &Repository{
		userCollection: &UserRepository{
			users: make(map[string]models.User),
		},
	}
}

func (r *Repository) GetUserTable() common.UserRepositoryInterface {
	return r.userCollection
}
