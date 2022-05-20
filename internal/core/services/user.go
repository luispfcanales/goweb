package services

import (
	"errors"

	"github.com/luispfcanales/goweb/internal/core/domain"
	"github.com/luispfcanales/goweb/internal/core/ports"
)

//UserService is template of service
type UserService struct {
	userRepository ports.Repository
}

//NewUserService return a userService
func NewUserService(repo ports.Repository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

//Create return new user
func (srv UserService) Create(name string) (domain.User, error) {
	user := domain.New("1234", name)
	u, err := srv.userRepository.Save(user)
	if err != nil {
		return domain.User{}, errors.New("user not created")
	}
	return u, nil
}
