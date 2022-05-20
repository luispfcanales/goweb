package domain

//User is template
type User struct {
	ID   string
	Name string
}

//New to return newUser
func New(id, name string) User {
	return User{
		ID:   id,
		Name: name,
	}
}
