package ports

import "github.com/luispfcanales/goweb/internal/core/domain"

//UserService is interface to services of users
type UserService interface {
	Create(name string) (domain.User, error)
}
