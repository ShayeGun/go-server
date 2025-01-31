package service

import "github.com/ShayeGun/go-server/internal/common"

type Service struct {
	userService common.UserServiceInterface
}

func NewService(cfg common.ExternalDependencies) (*Service, error) {
	us := NewUserService(cfg.GetUserTable())

	return &Service{
		userService: us,
	}, nil
}

func (sr *Service) GetUserService() common.UserServiceInterface {
	return sr.userService
}
