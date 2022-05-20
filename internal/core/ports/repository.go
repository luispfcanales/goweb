package ports

import "github.com/luispfcanales/goweb/internal/core/domain"

//Repository is a port
type Repository interface {
	Save(domain.User) (domain.User, error)
}
