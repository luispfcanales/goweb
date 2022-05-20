package repository

import "github.com/luispfcanales/goweb/internal/core/domain"

//Repository is template to data
type Repository struct {
	Data map[string]domain.User
}

//New return a new Repository
func New() *Repository {
	return &Repository{
		Data: map[string]domain.User{},
	}
}

//Save stores and return the user
func (r *Repository) Save(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
